package config

import (
	"MVC_DI/util"
	"reflect"
	"strconv"
	"strings"
)

// # Resolve
//
// Resolve resolves placeholders in the given struct.
//
// This function is useful for resolving placeholders in configuration structs.
// For example, if you have a struct like this:
//
//	type Config struct {
//		Database struct {
//			Host string `mapstructure:"host"`
//			Port int    `mapstructure:"port"`
//		} `mapstructure:"database"`
//	}
//
// And you want to set the `Host` field to the value of the `APP_HOST` field,
// you can do so like this:
//
// cfg := Config{}
//
// config.Resolve(&cfg)
//
// The `Resolve` function will replace placeholders in strings with actual values.
func Resolve[T any](v *T) {
	val := reflect.ValueOf(v).Elem()
	processValue(val, val, val.Type(), "")
}

// # processValue
//
// processValue processes the given value and replaces placeholders in strings
// with actual values from the configuration.
func processValue(root reflect.Value, val reflect.Value, typ reflect.Type, path string) {
	// Iterate over all fields of the given value.
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldPath := path + "." + fieldType.Name

		// If the field is a string, replace placeholders with actual values.
		if field.Kind() == reflect.String {
			field.SetString(replacePlaceholders(field.String(), root, val))
		}
		// If the field is a struct, recursively process its fields.
		if field.Kind() == reflect.Struct {
			processValue(root, field, fieldType.Type, fieldPath)
		}
	}
}

// # replacePlaceholders
//
// replacePlaceholders replaces placeholders in the given string with actual values
// from the configuration.
//
// The given string may contain placeholders in the form of ${path.to.field}.
// The function will replace these placeholders with the actual values from
// the given configuration struct.
//
// For example, if the configuration struct has a field `database.host` with
// value `localhost`, the string "postgres://${database.host}:5432/database" will
// be replaced with "postgres://localhost:5432/database".
func replacePlaceholders(s string, root reflect.Value, val reflect.Value) string {
	for {
		start := strings.Index(s, "${")
		if start == -1 {
			break
		}
		end := strings.Index(s[start:], "}") + start
		if end == -1 {
			break
		}

		placeholder := s[start+2 : end]
		replacement := resolvePlaceholder(placeholder, root, val)
		s = s[:start] + replacement + s[end+1:]
	}
	return s
}

// # resolvePlaceholder
//
// resolvePlaceholder resolves the given placeholder using the provided reflect values.
//
// The placeholder can be either a relative path (starting with a '.') or an absolute path.
// Relative paths are resolved against the provided value, while absolute paths are resolved
// against the root value.
func resolvePlaceholder(placeholder string, root reflect.Value, val reflect.Value) string {
	if strings.HasPrefix(placeholder, ".") {
		// If the placeholder is a relative path, resolve it using the current value.
		return resolveRelativePath(placeholder, val)
	}
	// Otherwise, resolve it as an absolute path using the root value.
	return resolveAbsolutePath(placeholder, root)
}

func resolveRelativePath(placeholder string, val reflect.Value) string {
	keys := strings.Split(placeholder[1:], ".")
	currentVal := val

	for _, key := range keys {
		key = util.SnakeToPascal(key)
		field, found := findField(currentVal, key)
		if !found {
			return ""
		}
		currentVal = field
	}

	return getValueAsString(currentVal)
}

func resolveAbsolutePath(placeholder string, val reflect.Value) string {
	keys := strings.Split(placeholder, ".")
	currentVal := val

	for _, key := range keys {
		key = util.SnakeToPascal(key)
		field, found := findField(currentVal, key)
		if !found {
			return ""
		}
		currentVal = field
	}

	return getValueAsString(currentVal)
}

func getValueAsString(val reflect.Value) string {
	switch val.Kind() {
	case reflect.String:
		return val.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	default:
		return ""
	}
}

func findField(val reflect.Value, key string) (reflect.Value, bool) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		if fieldType.Name == key {
			return field, true
		}
	}
	return reflect.Value{}, false
}
