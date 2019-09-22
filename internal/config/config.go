package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const envTagName = "env"

type Config struct {
	DiscordBotToken string `env:"DISCORD_BOT_TOKEN"`
	LogLevel        int    `env:"LOG_LEVEL"`
	Debug           bool   `env:"DEBUG"`
}

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
