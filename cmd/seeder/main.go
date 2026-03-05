package main

import (
	"log"

	_ "github.com/alf4ridzi/library-crud-ent-echo/ent/runtime"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database/seeder"
)

func main() {
	config.LoadEnv()

	client, err := database.NewMysqlEnt()
	if err != nil {
		log.Fatalf("failed create client : %v", client)
	}

	defer client.Close()

	seeder.Run(client)
}
