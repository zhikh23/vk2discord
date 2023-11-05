package main

import (
	"errors"
	"fmt"
	"vk2discord/internal/config"
	"vk2discord/internal/discordbot"
	"vk2discord/internal/postgres"
	"vk2discord/internal/vk2go"
)

func main() {
	cfg, err := config.Init(".env")
	if err != nil {
		panic(err)
	}
	db, err := postgres.Instance(cfg.Db)
	if err != nil {
		panic(err)
	}
	bot, err := discordbot.New(cfg.Discord)
	if err != nil {
		panic(err)
	}


	publications, err := vk2go.NewPublications(
		db,
		"its_bmstu",
		cfg.Vk.Token,
		cfg.Vk.ApiVer,
	)
	if err != nil {
		if errors.Is(err, vk2go.ErrNoNewPublications) {
			fmt.Print(err.Error())	
		} else {
			panic(err)
		}
	}


	const CHANNEL_ID = "963320221190471704"
	for _, publication := range publications {
		_, err := bot.Session.ChannelMessageSend(CHANNEL_ID, publication.Text)
		if err != nil {
			fmt.Printf("error during sent publication to discord: %s", err.Error())
		}
	}
}
