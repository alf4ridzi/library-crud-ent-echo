package seeder

import (
	"context"
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/role"
	"github.com/go-faker/faker/v4"
)

func UserSeeds(client *ent.Client) {
	userRole, err := client.Role.Query().Where(
		role.Name("user"),
	).First(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for range 10 {
		_, err := client.User.Create().
			SetEmail(faker.Email()).
			SetUsername(faker.Username()).
			SetName(faker.Name()).
			SetPassword("user12345").
			SetRole(userRole).
			Save(context.Background())

		if err != nil {
			log.Fatalf("error user seeeder : %v", err)
		}
	}
}
