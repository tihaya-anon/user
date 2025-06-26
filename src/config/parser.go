package config

import (
	"MVC_DI/global/module"
	"MVC_DI/util"
	"log"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func EnvParse[T any](file, env string, definition *T) error {
	file = file + "-" + env
	return Parse(file, definition)
}

// # Parse
//
// Parse Loads a config file into a struct
func Parse[T any](file string, definition *T) error {
	file = file + ".yaml"
	pathStr, fileStr := path.Split(file)
	pathStr = path.Join(module.GetResource(), pathStr)
	splitFile := strings.Split(fileStr, ".")

	name := splitFile[0]
	ext := splitFile[1]

	viper.SetConfigName(name)
	viper.SetConfigType(ext)
	viper.AddConfigPath(pathStr)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	decoderOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeHookFunc(time.RFC3339),
		mapstructure.StringToSliceHookFunc(","),
		func(f reflect.Type, t reflect.Type, data any) (any, error) {
			if f.Kind() == reflect.Map {
				newMap := make(map[string]any)
				for key, value := range data.(map[string]any) {
					newKey := util.SnakeToPascal(key)
					newMap[newKey] = value
				}
				return newMap, nil
			}
			return data, nil
		},
	))

	if err := viper.Unmarshal(definition, decoderOption); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return err
	}

	return nil
}
