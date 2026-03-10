package seeder

import (
	"context"
	"log"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
	"github.com/alf4ridzi/library-crud-ent-echo/ent/role"
)

var roles = []ent.Role{
	{
		Name: "admin",
	},
	{
		Name: "user",
	},
}

func RoleSeeder(client *ent.Client) {
	for _, r := range roles {
		exist, err := client.Role.Query().Where(
			role.Name(r.Name),
		).First(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			log.Fatal(err)
		}

		if exist != nil {
			continue
		}

		_, err = client.Role.
			Create().
			SetName(r.Name).
			Save(context.Background())

		if err != nil {
			log.Fatal(err)
		}
	}
}
