package test

import (
	"context"
	"log"
	"testing"

	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/infrastructure/database"
)

func TestConnectionHealthMysql(t *testing.T) {
	if err := config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Print("seems ok")
}

func TestQueryMysql(t *testing.T) {
	if err := config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

	db, err := database.NewMysqlEnt()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	users := db.User.Query().AllX(context.Background())
	t.Log(users)
}
