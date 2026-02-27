package seeder

import "github.com/alf4ridzi/library-crud-ent-echo/ent"

func Run(client *ent.Client) {
	UserSeeds(client)
}
