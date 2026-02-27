package migrate

import (
	"context"
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

func Run(client *ent.Client) {
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema : %v", err)
	}
}
