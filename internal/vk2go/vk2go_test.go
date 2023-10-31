package vk2go_test

import (
	"testing"
	"vk2discord/internal/config"
	"vk2discord/internal/vk2go"

	"github.com/sirupsen/logrus"
)

func TestGetDataFromVk(t *testing.T) {
	if err := config.Init("../../.env"); err != nil {
		t.Fatalf("error loading config: %s", err.Error())
	}
	
	msg, err := vk2go.GetDataFromVk(config.AppConfig.VkToken, "its_bmstu", 5.154)
	if err != nil {
		t.Fatalf("error getting data from VK: %s", err.Error())
	}
	
	logrus.Info(msg)
}
