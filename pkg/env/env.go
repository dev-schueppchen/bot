package env

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func ParseEnv(v interface{}) error {
	tv := reflect.ValueOf(v)

	for i := 0; i < tv.NumField(); i++ {
		fieldType := tv.Field(i)
		tag := fieldType.Elem().Tag.Get(envTagName)
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

	return nil, nil

	json.Unmarshal()
}
