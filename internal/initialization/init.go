package initialization

import (
	"github.com/andersfylling/disgord"
	"github.com/dev-schueppchen/bot/internal/config"
	"github.com/dev-schueppchen/bot/internal/handler"
	"github.com/dev-schueppchen/bot/internal/logger"
)

// Init initializes life time handler necessary for
// all bot functions.
//
// This function blocks the current go routine and
// to enther the Disgord event loop.
func Init() {

	// ---------------------------------
	// Logger initialization

	logger.Setup(`%{color}▶  %{level:.4s} %{id:03d}%{color:reset} %{message}`, 5)
	mlog := logger.GetLogger()

	// ---------------------------------
	// Config initialization

	cfg, err := config.ReadFromEnv()
	if err != nil {
		mlog.Fatalf("failed reading config from environment: %s", err.Error())
	}

	logger.SetLogLevel(cfg.LogLevel)

	// ---------------------------------
	// Disgord initialization

	dc := disgord.New(&disgord.Config{
		BotToken: cfg.DiscordBotToken,
	})

	defer dc.StayConnectedUntilInterrupted()

	dc.Ready(handler.NewReady(dc).Handler)
}
