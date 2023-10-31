package db

import (
	"context"
	"fmt"
	"vk2discord/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Db struct {
	Pool *pgxpool.Pool
}

var db Db

func Init() error {
	dburl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.AppConfig.Db.User,
		config.AppConfig.Db.Password,
		config.AppConfig.Db.Host,
		config.AppConfig.Db.Port,
		config.AppConfig.Db.Name,
	)
	logrus.Infof("Using db url: %s", dburl)
	pool, err := pgxpool.Connect(context.Background(), dburl)
	if err != nil {
		logrus.Errorf("Failed to connect to db: %v", err)
		return err
	}
	logrus.Infof("Successfully connected to database")

	db = Db{Pool: pool}
	return nil
}

func Close() {
	db.Pool.Close()
}

func NewConn() (*pgxpool.Conn, error) {
	return db.Pool.Acquire(context.Background())
}
