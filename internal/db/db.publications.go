package db

import (
	"context"

	"github.com/sirupsen/logrus"
)

func IsPublished(domain string, id int) (bool, error) {
	conn, err := NewConn()
	if err != nil {
		return false, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"SELECT 1 FROM publications WHERE domain=$1 AND id=$2",
		domain, id)
	var publicated int
	if err = row.Scan(&publicated); err != nil {
		return false, nil
	}
	return true, nil
}

func MarkAsPublished(domain string, id int) error {
	conn, err := NewConn()
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(),
		"INSERT INTO publications (domain, id) VALUES ($1, $2)",
		domain, id)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		logrus.Error(err.Error())
	}
		
	return nil
}
