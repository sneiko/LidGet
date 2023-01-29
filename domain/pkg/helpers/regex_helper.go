package helpers

import (
	"lidget/domain/entities"
	"lidget/domain/pkg/logger"
	"regexp"
)

func CheckStringForPattern(row string, patterns []entities.CategoryPattern) []entities.CategoryPattern {
	var matchesPatterns []entities.CategoryPattern
	for _, pattern := range patterns {
		isMatch, err := regexp.MatchString(pattern.Pattern, row)
		if err != nil {
			logger.Error(err)
		}

		if isMatch {
			matchesPatterns = append(matchesPatterns, pattern)
		}
	}
	return matchesPatterns
}
