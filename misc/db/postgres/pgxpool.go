// Package postgres contains functions related to connecting and interacting with a PostgreSQL database using the pgx library.
package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PoolIface interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewPool(pool *pgxpool.Pool) *Pool {
	return &Pool{pool: pool}
}

type Pool struct {
	pool *pgxpool.Pool
}

func (o *Pool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return o.pool.QueryRow(ctx, sql, args)
}

func (o *Pool) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return o.pool.BeginTx(ctx, txOptions)
}

// NewPGXPool creates a new connection pool to a PostgreSQL database using the pgx library.
func NewPGXPool(ctx context.Context, uri string, simpleProtocol bool) (*pgxpool.Pool, error) {
	var (
		err    error
		conn   *pgxpool.Pool
		pgconf *pgxpool.Config
	)

	pgconf, err = pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}

	pgconf.ConnConfig.PreferSimpleProtocol = simpleProtocol

	if conn, err = pgxpool.ConnectConfig(ctx, pgconf); err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return conn, nil
}
