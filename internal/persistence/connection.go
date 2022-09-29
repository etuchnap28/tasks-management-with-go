package persistence

import (
	"embed"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/config"
	"github.com/gobuffalo/pop/v6"
)

//go:embed migrations/*.fizz
var migrations embed.FS

func OpenDBConnection(cfg *config.DatabaseConfig) *pop.Connection {
	var dbConn *pop.Connection
	var migrationBox pop.MigrationBox
	var err error
	if dbConn, err = pop.NewConnection(
		&pop.ConnectionDetails{
			Dialect:  cfg.Dialect,
			URL:      cfg.URL,
			Database: cfg.Database,
			Host:     cfg.Host,
			Port:     cfg.Port,
			User:     cfg.User,
			Password: cfg.Password,
			Pool:     cfg.Pool,
			IdlePool: cfg.IdlePool,
		},
	); err != nil {
		panic(err)
	}
	if err = dbConn.Open(); err != nil {
		panic(err)
	}
	if migrationBox, err = pop.NewMigrationBox(migrations, dbConn); err != nil {
		panic(err)
	}
	if err = migrationBox.Up(); err != nil {
		panic(err)
	}
	return dbConn
}
