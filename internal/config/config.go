package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	VkToken string
}

func Init(fromFile string) error {
	err := godotenv.Load(fromFile)
	if err != nil{
		return err
	}	

	vkToken := os.Getenv("VK_TOKEN")
	if vkToken == "" {
		return fmt.Errorf("missing required VK_TOKEN in .env file")
	}

	AppConfig = Config{
		VkToken: vkToken,
	}
	return nil
}
