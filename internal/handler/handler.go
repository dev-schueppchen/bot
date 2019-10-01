package handler

import (
	"github.com/andersfylling/disgord"
	"github.com/dev-schueppchen/bot/internal/logger"
)

// Ready event handler.
// Fires when all shard insatnces of the bot are ready.
type Ready struct {
	instance *disgord.Client
}

// NewReady creates an new instance of Ready event handler
//
// Parameters
//   instance : Disgord client instance
//
// Returns
//   Ready handler instance
func NewReady(instance *disgord.Client) *Ready {
	return &Ready{
		instance: instance,
	}
}

// Handler is the function which is fired by the
// Disgord instance when all shards are ready.
func (h *Ready) Handler() {
	selfUser, err := h.instance.GetCurrentUser()

	if err != nil {
		logger.GetLogger().Errorf("Getting self user failed: %s", err.Error())
		return
	}

	logger.GetLogger().Infof("Ready as user %s#%s (%s)",
		selfUser.Username, selfUser.Discriminator, selfUser.ID)
}
