package initialization

import (
	"github.com/andersfylling/disgord"
	"github.com/dev-schueppchen/bot/internal/config"
	"github.com/dev-schueppchen/bot/internal/logger"
)

func Init() {
	logger.Setup(`%{color}▶  %{level:.4s} %{id:03d}%{color:reset} %{message}`, 5)
	mlog := logger.GetLogger()

	cfg, err := config.ReadFromEnv()
	if err != nil {
		mlog.Fatalf("failed reading config from environment: %s", err.Error())
	}

	logger.SetLogLevel(cfg.LogLevel)

	dc := disgord.New(&disgord.Config{
		BotToken: cfg.DiscordBotToken,
	})

	defer dc.StayConnectedUntilInterrupted()

	dc.Ready(func() {
		logger.GetLogger().Info("READY")
	})
}
