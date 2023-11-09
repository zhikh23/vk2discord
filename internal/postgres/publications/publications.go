package publications

import (
	"context"
	"vk2discord/internal/config"
	"vk2discord/internal/postgres"

	"github.com/sirupsen/logrus"
)


type PostgresPublicationStore struct {
	db *postgres.Db
}


func NewStore(cfg *config.DbConfig) (*PostgresPublicationStore, error) {
	db, err := postgres.Init(cfg)
	if err != nil {
		return nil, err
	}
	return &PostgresPublicationStore{
		db: db,
	}, nil
}


func (store *PostgresPublicationStore) Close() {
	store.db.Close()
}


func (store *PostgresPublicationStore) IsPublished(pubId int, domain string) (bool, error) {
	conn, err := store.db.NewConn()
	if err != nil {
		return false, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"SELECT 1 FROM publications WHERE domain=$1 AND id=$2",
		domain, pubId)
	var publicated int
	if err = row.Scan(&publicated); err != nil {
		return false, nil
	}
	return true, nil
}


func (store *PostgresPublicationStore) MarkAsPublished(pubId int, domain string) error {
	conn, err := store.db.NewConn()
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"INSERT INTO publications (domain, id) VALUES ($1, $2)",
		domain, pubId)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		logrus.Error(err.Error())
	}
		
	return nil
}


func (store * PostgresPublicationStore) RemoveFromPublished(pubId int, domain string) error {
	conn, err := store.db.NewConn()
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"DELETE FROM publications WHERE domain=$1 AND id=$2",
		domain, pubId)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		logrus.Error(err.Error())
	}
		
	return nil
}
