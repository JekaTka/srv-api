package pg

import (
	"database/sql"
	"context"

	// blank import to have driver registered
	_ "github.com/lib/pq"
)

type Client interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	EnsureConnected(ctx context.Context) error
	Ping(ctx context.Context) error
}

type client struct {
	db *sql.DB
}

func New(/*config Config*/) Client {
	return &client{/* config.db */}
}
