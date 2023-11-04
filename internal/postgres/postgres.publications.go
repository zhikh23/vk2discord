package postgres

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (db *Db) IsPublished(pubId int, domain string) (bool, error) {
	conn, err := db.NewConn()
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

func (db *Db) MarkAsPublished(pubId int, domain string) error {
	conn, err := db.NewConn()
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
