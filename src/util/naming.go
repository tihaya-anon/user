package util

import (
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