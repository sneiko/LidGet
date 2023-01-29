package tgClient

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/telegram/updates/hook"
	tg "github.com/gotd/td/tg"
	"github.com/jackc/pgx/v5/pgxpool"
	"lidget/domain/config"
	"lidget/domain/entities"
	"lidget/domain/pkg/helpers"
	"lidget/domain/pkg/logger"
	"lidget/domain/repository"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type TgClient struct {
	db           *pgxpool.Pool
	repositories repository.Repositories

	patterns []entities.CategoryPattern

	appId       int
	appHash     string
	phoneNumber string
	password    string
}

func NewTgClient(cfg *config.TgConfig, db *pgxpool.Pool) *TgClient {
	return &TgClient{
		db:           db,
		appId:        cfg.Id,
		appHash:      cfg.Hash,
		phoneNumber:  cfg.PhoneNumber,
		password:     cfg.Password,
		repositories: repository.NewRepositories(db),
	}
}

func (s *TgClient) fetchData(ctx context.Context) error {
	patterns, err := s.repositories.CategoryPatterns.GetAll(ctx)
	if err != nil {
		return err
	}

	s.patterns = patterns

	return nil
}

func (s *TgClient) Run(ctx context.Context) error {
	if err := s.fetchData(ctx); err != nil {
		return err
	}

	d := tg.NewUpdateDispatcher()
	gaps := updates.New(updates.Config{
		Handler: d,
	})

	path, _ := os.Executable()
	client := telegram.NewClient(s.appId, s.appHash,
		telegram.Options{
			UpdateHandler: gaps,
			Middlewares: []telegram.Middleware{
				hook.UpdateHook(gaps.Handle),
			},
			SessionStorage: &telegram.FileSessionStorage{
				Path: filepath.Dir(path) + "/tg_session",
			},
		})

	codePrompt := func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
		fmt.Print("Enter code: ")
		code, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(code), nil
	}

	d.OnNewChannelMessage(s.newChannelMessage)
	d.OnNewMessage(s.newMessage)

	var authenticator auth.UserAuthenticator
	if s.password == "" {
		authenticator = auth.CodeOnly(s.phoneNumber, auth.CodeAuthenticatorFunc(codePrompt))
	} else {
		authenticator = auth.Constant(s.phoneNumber, s.password, auth.CodeAuthenticatorFunc(codePrompt))
	}
	flow := auth.NewFlow(
		authenticator,
		auth.SendCodeOptions{},
	)

	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			logger.Info(err)
			return err
		}

		user, err := client.Self(ctx)
		if err != nil {
			return err
		}

		if err := gaps.Auth(ctx, client.API(), user.ID, true, false); err != nil {
			return err
		}
		defer func() { _ = gaps.Logout() }()

		logger.Info("Running!")

		<-ctx.Done()
		return ctx.Err()
	})
}

func (s *TgClient) newChannelMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
	if msg, ok := update.Message.(*tg.Message); ok {
		if msg.Message != "" && msg.FromID != nil {
			userId := msg.FromID.(*tg.PeerUser).UserID
			username := e.Users[userId].Username

			if username == "" {
				s.handleMsg(ctx, msg.Message, strconv.Itoa(int(userId)))
			} else {
				isBot, err := regexp.MatchString("bot", username)
				if err != nil {
					logger.Error(err)
					return nil
				}

				if !isBot {
					s.handleMsg(ctx, msg.Message, fmt.Sprintf("@%s", username))
				}
			}
		}
	}
	return nil
}

func (s *TgClient) newMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
	if msg, ok := update.Message.(*tg.Message); ok {
		if msg.Message != "" && msg.FromID != nil {
			userId := msg.FromID.(*tg.PeerUser).UserID
			s.handleMsg(ctx, msg.Message, strconv.Itoa(int(userId)))
		}
	}
	return nil
}

func (s *TgClient) handleMsg(ctx context.Context, msg string, fromId string) {
	matches := helpers.CheckStringForPattern(msg, s.patterns)
	if len(matches) > 0 {
		logger.Info(fmt.Sprintf("fromId: %s; msg: %s", fromId, msg))
		if !s.repositories.Requests.IsExistByText(ctx, msg) {
			row := entities.Request{
				SenderTgId: fromId,
				Text:       msg,
			}

			requestId, err := row.Create(ctx, s.db)
			if err != nil {
				logger.Error(err)
			}

			for _, m := range matches {
				requestCategory := entities.RequestCategory{
					RequestId:  *requestId,
					CategoryId: m.CategoryId,
				}

				_, err = requestCategory.Create(ctx, s.db)
				if err != nil {
					logger.Error(err)
				}
			}
		}
	}
}
