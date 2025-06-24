package util

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func SnakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	caser := cases.Title(language.English)
	for i, part := range parts {
		parts[i] = caser.String(part)
	}
	return strings.Join(parts, "")
}

func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	caser := cases.Title(language.English)
	for i, part := range parts {
		if i == 0 {
			continue
		}
		parts[i] = caser.String(part)
	}
	return strings.Join(parts, "")
}

func SnakeToHyphen(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

func PascalToSnake(s string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(s, `${1}_${2}`)
	return strings.ToLower(snake)
}
