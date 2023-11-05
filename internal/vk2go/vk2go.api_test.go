package vk2go

import (
	"testing"
	"vk2discord/internal/config"

	"github.com/sirupsen/logrus"
)

func TestGetDataFromVk(t *testing.T) {
	if err := config.Init("../../.env"); err != nil {
		t.Fatalf("error loading config: %s", err.Error())
	}
	
	res, err := fetchDataFromVk("its_bmstu", 5, config.AppConfig.Vk.Token, config.AppConfig.Vk.ApiVer)
	if err != nil {
		t.Fatalf("error getting data from VK: %s", err.Error())
	}
	
	msg := res.Response.Items[0].Text
	if msg == "" {
		t.Fatal("returned text must be not empty")
	}
	logrus.Info(msg)
}
