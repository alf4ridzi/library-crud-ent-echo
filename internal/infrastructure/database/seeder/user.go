package seeder

import (
	"context"
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/internal/pkg/cryptoutil"
	"github.com/go-faker/faker/v4"
)

func UserSeeds(client *ent.Client) {
	for range 10 {
		_, err := client.User.Create().
			SetEmail(faker.Email()).
			SetUsername(faker.Username()).
			SetName(faker.Name()).
			SetPassword(cryptoutil.HashPassword("user12345")).
			Save(context.Background())

		if err != nil {
			log.Fatalf("error user seeeder : %v", err)
		}
	}
}
