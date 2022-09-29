package main

import (
	"github.com/etuchnap28/tasks-mangement-with-go/internal/api"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/config"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/persistence"
	"github.com/gobuffalo/pop/v6"
	"log"
)

func main() {
	config := config.ReadAppConfig()
	dbConn := persistence.OpenDBConnection(config.Database)
	defer dbConn.Close()
	runGinServer(config.Server, dbConn)
}

func runGinServer(config *config.ServerConfig, store *pop.Connection) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.Address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
