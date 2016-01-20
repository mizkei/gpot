package gpot

import (
	"golang.org/x/oauth2"
	plusd "google.golang.org/api/plusdomains/v1"
)

// BotPlus is simple bot.
type BotPlus struct {
	service plusd.Service
}

// Post post comment
func (b *BotPlus) Post() error {
	return nil
}

// EventTicker returns channel that send empty
// when activity posted.
func (b *BotPlus) EventTicker() <-chan struct{} {
	ch := make(chan struct{})

	return ch
}

// NewBotPlus generate BotPlus.
func NewBotPlus(config oauth2.Config) *BotPlus {
	return nil
}
