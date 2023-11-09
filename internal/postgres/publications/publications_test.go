package publications_test

import (
	"testing"
	"vk2discord/internal/config"
	"vk2discord/internal/postgres/publications"
)

const (
	testPublicationId = 2
	testDomain = "its_bmstu"
)


func TestPublications(t *testing.T) {
	cfg, err := config.Init("../../../.env")
	if err != nil {
		t.Fatalf("error load config: %s", err.Error())
	}
	
	store, err := publications.NewStore(&cfg.Db)
	if err != nil {
		t.Fatalf("error while initialize storage: %s", err.Error())
	}
	defer store.Close()

	t.Run("testing store.MarkAsPublished", func(t *testing.T) {
		err := store.MarkAsPublished(testPublicationId, testDomain)
		if err != nil {
			t.Fatalf("store.MarkAsPublished returned error: %s", err.Error())
		}
	})

	t.Run("testing store.IsPublished", func(t *testing.T) {
		published, err := store.IsPublished(testPublicationId, testDomain)
		if err != nil {
			t.Fatalf("store.IsPublished returned error: %s", err.Error())
		}
		if !published {
			t.Fatalf("store.IsPublished must return true")
		}
	})

	t.Run("testing store.RemoveFromPublished", func(t *testing.T) {
		err := store.RemoveFromPublished(testPublicationId, testDomain)
		if err != nil {
			t.Fatalf("store.RemoveFromPublished returned error: %s", err.Error())
		}
	})
}
