package discordbot

import (
	"vk2discord/internal/config"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
}

func New(cfg config.DiscordConfig) (*Bot, error) {
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		Session: session,
	}, nil
}
