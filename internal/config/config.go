package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const envTagName = "env"

// Config describes the general configuration
// vaules which are set by environment variables.
type Config struct {
	DiscordBotToken string `env:"DISCORD_BOT_TOKEN"`
	LogLevel        int    `env:"LOG_LEVEL"`
	Debug           bool   `env:"DEBUG"`
}

// ReadFromEnv reads the set environment variable
// keys as described in Config and returns the
// Config instance containing the read values.
//
// The keys are collected by reflecting the given
// struct tags. The key 'env' defines the key which
// is read from the environemnt and the corresponding
// field will be set by its value.
//
// Currently supported field types are
//  - string
//  - int
//  - bool
// If you need another type, add it to the switch
// below by defininf how to parse the string value
// read from environemt to the target value type.
// Every undefined value type will return an
// 'unsupported type' error.
func ReadFromEnv() (*Config, error) {
	cfg := Config{}
	tv := reflect.ValueOf(&cfg)
	tt := reflect.TypeOf(cfg)

	for i := 0; i < tt.NumField(); i++ {
		fieldType := tt.Field(i)
		tag := fieldType.Tag.Get(envTagName)
		val := os.Getenv(tag)

		if val == "" {
			continue
		}

		field := tv.Elem().Field(i)

		switch field.Kind() {

		case reflect.String:
			field.SetString(val)

		case reflect.Int:
			vi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			field.SetInt(int64(vi))

		case reflect.Bool:
			vb := val == "1" || strings.ToLower(val) == "true"
			field.SetBool(vb)

		default:
			return nil, fmt.Errorf("unsupported type")
		}
	}

	return &cfg, nil
}
