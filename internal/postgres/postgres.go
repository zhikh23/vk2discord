package postgres

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

func Init(cfg *config.DbConfig) (*Db, error) {
	dburl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	logrus.Infof("Using db url: %s", dburl)
	pool, err := pgxpool.Connect(context.Background(), dburl)
	if err != nil {
		logrus.Errorf("Failed to connect to db: %v", err)
		return nil, err
	}
	logrus.Infof("Successfully connected to database")

	db := Db{Pool: pool}
	return &db, err
}

func (db *Db) Close() {
	db.Pool.Close()
}

func (db *Db) NewConn() (*pgxpool.Conn, error) {
	return db.Pool.Acquire(context.Background())
}
