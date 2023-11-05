package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Vk      VkConfig
	Db      DbConfig
	Discord DiscordConfig
}

type VkConfig struct {
	Token  string
	ApiVer float32
}

type DbConfig struct {
	Host 	 string
	Port 	 int
	Name 	 string
	User 	 string
	Password string
}

type DiscordConfig struct {
	Token string
}

func Init(fromFile string) (*Config, error) {
	err := godotenv.Load(fromFile)
	if err != nil{
		return nil, err
	}	

	vkToken := os.Getenv("VK_TOKEN")
	if vkToken == "" {
		return nil, fmt.Errorf("missing required VK_TOKEN in .env file")
	}

	vkApiVerStr := os.Getenv("VK_API_VERSION")
	if vkApiVerStr == "" {
		return nil, fmt.Errorf("missing required VK_API_VERSION in .env file")
	}
	vkApiVer, err := strconv.ParseFloat(vkApiVerStr, 32)
	if err != nil {
		return nil, fmt.Errorf("error during parsing VK_API_VERSION: %s", err.Error())
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := 5432
	strPort := os.Getenv("DB_PORT")
	if strPort != "" {
		v, err := strconv.Atoi(strPort)
		if err != nil {
			port = v
		}
	}

	name := os.Getenv("DB_NAME")
	if host == "" {
		return nil, fmt.Errorf("missing required DB_NAME in .env file")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, fmt.Errorf("missing required DB_NAME in .env file")
	}

	password := os.Getenv("DB_PASSWORD")

	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		return nil, fmt.Errorf("missing reqired DISCORD_TOKEN in .env file")
	}

	return &Config{
		Vk: VkConfig{
			Token: vkToken,
			ApiVer: float32(vkApiVer),
		},
		Db: DbConfig {
			Host: host,
			Port: port,
			Name: name,
			User: user,
			Password: password,
		},
		Discord: DiscordConfig{
			Token: discordToken,
		},
	}, nil
}
