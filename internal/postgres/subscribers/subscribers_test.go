package subscribers_test

import (
	"testing"
	"vk2discord/internal/config"
	"vk2discord/internal/discordbot"
	"vk2discord/internal/postgres/subscribers"
)

const (
	testChannelId = discordbot.ChannelId(963320221190471704)
	testDomain    = "its_bmstu"
)


func TestSubscribers(t *testing.T) {
	cfg, err := config.Init("../../../.env")
	if err != nil {
		t.Fatalf("error load config: %s", err.Error())
	}
	
	store, err := subscribers.NewStore(&cfg.Db)
	if err != nil {
		t.Fatalf("error while initialize storage: %s", err.Error())
	}
	defer store.Close()

	t.Run("testing store.NewSubscriber", func(t *testing.T) {
		err := store.NewSubsciber(testDomain, testChannelId)
		if err != nil {
			t.Fatalf("store.NewSubscriber returned error: %s", err.Error())
		}
	})

	t.Run("testing store.SubscribedFor", func (t *testing.T)  {
		subscribers, err := store.SubscribedFor(testDomain)
		if err != nil {
			t.Fatalf("store.SubscribedFor returned error: %s", err.Error())
		}
		subscribed := false
		for _, channelId := range subscribers {
			if channelId == testChannelId {
				subscribed = true
				break
			}
		}
		if !subscribed {
			t.Fatalf("store.SubscribedFor must return channelId=%d", testChannelId)
		}
	})
	
	t.Run("testing db.Ubsubscribe", func (t *testing.T)  {
		err = store.Unsubscribe(testDomain, testChannelId)
		if err != nil {
			t.Fatalf("store.Umsubscribe returned error: %s", err.Error())
		}
	})
}
