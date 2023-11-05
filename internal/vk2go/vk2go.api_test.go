package vk2go

import (
	"testing"
	"vk2discord/internal/config"

	"github.com/sirupsen/logrus"
)

func TestGetDataFromVk(t *testing.T) {
	cfg, err := config.Init("../../.env") 
	if err != nil {
		t.Fatalf("error loading config: %s", err.Error())
	}
	
	res, err := fetchDataFromVk("its_bmstu", 5, cfg.Vk.Token, cfg.Vk.ApiVer)
	if err != nil {
		t.Fatalf("error getting data from VK: %s", err.Error())
	}
	
	msg := res.Response.Items[0].Text
	if msg == "" {
		t.Fatal("returned text must be not empty")
	}
	logrus.Info(msg)
}
