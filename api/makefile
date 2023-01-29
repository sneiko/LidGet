
build_api:
	GOOS='linux' CGO_ENABLED=0 go build -ldflags "-s -w" -o ./build/api ./cmd/api/main.go
	upx ./build/api