package storage

import (
	"context"
	"github.com/go-list-templ/grpc/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	*pgxpool.Pool
}

func NewPostgres(cfg *config.DB) (*Postgres, error) {
	conf, err := pgxpool.ParseConfig(cfg.Url)
	if err != nil {
		return nil, err
	}

	conf.MaxConns = int32(cfg.MaxConn)
	conf.MaxConnLifetime = cfg.MaxConnTime

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &Postgres{pool}, nil
}
