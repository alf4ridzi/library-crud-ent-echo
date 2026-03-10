package seeder

import "github.com/alf4ridzi/library-crud-ent-echo/ent"

func Run(client *ent.Client) {
	RoleSeeder(client)
	UserSeeds(client)
	CategorySeed(client)
	BookSeeder(client)
}
