package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	VkToken string
	Db DbConfig
}

type DbConfig struct {
	Host string
	Port int
	Name string
	User string
	Password string
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
		return fmt.Errorf("missing required DB_NAME in .env file")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return fmt.Errorf("missing required DB_NAME in .env file")
	}

	password := os.Getenv("DB_PASSWORD")

	AppConfig = Config{
		VkToken: vkToken,
		Db: DbConfig {
			Host: host,
			Port: port,
			Name: name,
			User: user,
			Password: password,
		},
	}
	return nil
}
