package main

import (
	"errors"
	"fmt"
	"vk2discord/internal/config"
	"vk2discord/internal/postgres"
	"vk2discord/internal/vk2go"
)

func main() {
	err := config.Init(".env")
	if err != nil {
		panic(err)
	}

	db, err := postgres.Instance()
	if err != nil {
		panic(err)
	}

	publications, err := vk2go.NewPublications(
		db,
		"its_bmstu",
		config.AppConfig.VkToken,
		config.AppConfig.VkApiVer,
	)
	if errors.Is(err, vk2go.ErrNoNewPublications) {
		fmt.Print(err.Error())	
	} else {
		panic(err)
	}

	for _, publication := range publications {
		fmt.Printf("%v\n", publication)
	}
}
