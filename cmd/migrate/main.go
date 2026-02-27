package main

import (
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database/migrate"
)

func main() {
	config.LoadEnv()

	client, err := database.NewMysqlEnt()
	if err != nil {
		log.Fatalf("failed create client : %v", client)
	}

	migrate.Run(client)
}
