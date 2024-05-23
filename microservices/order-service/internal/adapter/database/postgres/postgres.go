package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // importa o drive do postgres para rodar as migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

type DatabaseClientConfig struct {
	databaseURL     string
	context         context.Context
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
	runMigrations   bool
}

type DatabaseOption func(d *DatabaseClientConfig)

func Context(ctx context.Context) DatabaseOption {
	return func(c *DatabaseClientConfig) {
		c.context = ctx
	}
}

func DatabaseURL(url string) DatabaseOption {
	return func(c *DatabaseClientConfig) {
		c.databaseURL = url
	}
}

func SetMaxOpenConns(value int) DatabaseOption {
	return func(c *DatabaseClientConfig) {
		c.maxOpenConns = value
	}
}

func SetMaxIdleConns(value int) DatabaseOption {
	return func(c *DatabaseClientConfig) {
		c.maxIdleConns = value
	}
}

func SetConnMaxLifetime(value time.Duration) DatabaseOption {
	return func(c *DatabaseClientConfig) {
		c.connMaxLifetime = value
	}
}

func RunMigrations(value bool) DatabaseOption {
	return func(c *DatabaseClientConfig) {
		c.runMigrations = value
	}
}

// Initialize retorna um pool de conexões com o banco de dados
func Initialize(opts ...DatabaseOption) *sqlx.DB {
	databaseOptions := &DatabaseClientConfig{
		maxOpenConns:    100,
		maxIdleConns:    10,
		connMaxLifetime: 0,
		runMigrations:   true,
	}
	for _, opt := range opts {
		opt(databaseOptions)
	}

	db, err := sqlx.ConnectContext(databaseOptions.context, "nrpostgres", databaseOptions.databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.DB.SetMaxOpenConns(databaseOptions.maxOpenConns)
	db.DB.SetMaxIdleConns(databaseOptions.maxIdleConns)
	db.DB.SetConnMaxLifetime(databaseOptions.connMaxLifetime)

	if databaseOptions.runMigrations {
		runMigrations(databaseOptions.databaseURL)
	}

	return db
}

func runMigrations(databaseURL string) {
	m, err := migrate.New("file://database/migrations", databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
