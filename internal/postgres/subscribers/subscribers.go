package subscribers

import (
	"context"
	"errors"
	"vk2discord/internal/config"
	"vk2discord/internal/discordbot"
	"vk2discord/internal/postgres"

	"github.com/sirupsen/logrus"
)


var (
	ErrNoSubscribersFound = errors.New("no subscribers found")
)


type PostgresSubscribersStore struct {
	db *postgres.Db
}


func NewStore(cfg *config.DbConfig) (*PostgresSubscribersStore, error) {
	db, err := postgres.Init(cfg)
	if err != nil {
		return nil, err
	}
	return &PostgresSubscribersStore{
		db: db,
	}, nil
}


func (store *PostgresSubscribersStore) Close() {
	store.db.Close()
}


func (store *PostgresSubscribersStore) NewSubsciber(domain string, channelId discordbot.ChannelId) (error) {
	conn, err := store.db.NewConn()
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"INSERT INTO subscribers (domain, channel_id) VALUES ($1, $2)",
		domain, channelId)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		logrus.Error(err.Error())
	}
		
	return nil
}


func (store *PostgresSubscribersStore) SubscribedFor(domain string) ([]discordbot.ChannelId, error) {
	conn, err := store.db.NewConn()
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		"SELECT channel_id FROM subscribers WHERE domain=$1", domain)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channelIds []discordbot.ChannelId
	var channelId discordbot.ChannelId
	for rows.Next() {
		err := rows.Scan(&channelId)
		if err != nil {
			logrus.Errorf("error during fetching subscribers: %s", err.Error())
		} else {
			channelIds = append(channelIds, channelId)
		}
	}
	if len(channelIds) == 0 {
		return nil, ErrNoSubscribersFound
	}
	return channelIds, nil
}


func (store *PostgresSubscribersStore) Unsubscribe(domain string, channelId discordbot.ChannelId) (error) {
	conn, err := store.db.NewConn()
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"DELETE FROM subscribers WHERE domain=$1 AND channel_id=$2",
		domain, channelId)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		logrus.Error(err.Error())
	}
		
	return nil
}
