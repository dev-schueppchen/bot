package main

import (
	"fmt"

	"github.com/dev-schueppchen/bot/internal/config"
	"github.com/dev-schueppchen/bot/internal/logger"
)

func main() {
	logger.Setup(`%{color}▶  %{level:.4s} %{id:03d}%{color:reset} %{message}`, 5)
	mlog := logger.GetLogger()

	cfg, err := config.ReadFromEnv()
	if err != nil {
		mlog.Fatalf("failed reading config from environment: %s", err.Error())
	}

	logger.SetLogLevel(cfg.LogLevel)

	fmt.Printf("%+v\n", cfg)
}
