package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/config"
)

func NewMysqlEnt() (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.AppConfig.DBUserName,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
