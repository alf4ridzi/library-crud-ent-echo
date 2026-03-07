package seeder

import (
	"context"
	"log"
	"time"

	"github.com/alf4ridzi/library-crud-ent-echo/ent"
)

var books = []ent.Books{
	{
		Author:            "Jack Hanma",
		Description:       "Ini adalah buku fiksi tentang boneka tua di rumah terbengkalai",
		Title:             "Old Doll",
		Quantity:          10,
		AvailableQuantity: 10,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 1}, // Fiction
				{ID: 6}, // Biography
				{ID: 8}, // Mystery
			},
		},
	},
	{
		Author:            "Sarah Connor",
		Description:       "Kisah nyata seorang ilmuwan yang berjuang menemukan vaksin di tengah wabah global",
		Title:             "The Last Cure",
		Quantity:          15,
		AvailableQuantity: 15,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 2}, // Non-Fiction
				{ID: 3}, // Science
				{ID: 6}, // Biography
			},
		},
	},
	{
		Author:            "Reza Mahendra",
		Description:       "Panduan lengkap membangun aplikasi modern menggunakan Go dan microservices",
		Title:             "Go Beyond: Modern App Development",
		Quantity:          20,
		AvailableQuantity: 20,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 4}, // Technology
				{ID: 2}, // Non-Fiction
			},
		},
	},
	{
		Author:            "Emily Hartwell",
		Description:       "Petualangan seorang penyihir muda yang mencari pedang legenda di dunia paralel",
		Title:             "Sword of Aeloria",
		Quantity:          12,
		AvailableQuantity: 12,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 7}, // Fantasy
				{ID: 1}, // Fiction
			},
		},
	},
	{
		Author:            "Dr. Ahmad Fauzi",
		Description:       "Menelusuri peristiwa besar Perang Dunia II dari sudut pandang Asia Tenggara",
		Title:             "Asia in the Storm",
		Quantity:          8,
		AvailableQuantity: 8,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 5}, // History
				{ID: 2}, // Non-Fiction
			},
		},
	},
	{
		Author:            "Laura Simmons",
		Description:       "Detektif muda memecahkan serangkaian pembunuhan misterius di kota kecil bersalju",
		Title:             "Frozen Secrets",
		Quantity:          18,
		AvailableQuantity: 18,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 8}, // Mystery
				{ID: 1}, // Fiction
			},
		},
	},
	{
		Author:            "Budi Santoso",
		Description:       "Teknik-teknik praktis untuk meningkatkan produktivitas dan mencapai tujuan hidup",
		Title:             "Hidup Produktif Setiap Hari",
		Quantity:          25,
		AvailableQuantity: 25,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 9}, // Self-Help
				{ID: 2}, // Non-Fiction
			},
		},
	},
	{
		Author:            "Nina Karlova",
		Description:       "Biografi lengkap Marie Curie, ilmuwan perempuan pertama peraih Nobel dua kali",
		Title:             "Radiant: The Life of Marie Curie",
		Quantity:          10,
		AvailableQuantity: 10,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 6}, // Biography
				{ID: 3}, // Science
				{ID: 5}, // History
			},
		},
	},
	{
		Author:            "Tommy Wijaya",
		Description:       "Petualangan si Kancil dan teman-temannya menjelajahi hutan ajaib penuh kejutan",
		Title:             "Kancil dan Hutan Ajaib",
		Quantity:          30,
		AvailableQuantity: 30,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 10}, // Children
				{ID: 7},  // Fantasy
			},
		},
	},
	{
		Author:            "Marcus Ellison",
		Description:       "Eksplorasi mendalam tentang bagaimana kecerdasan buatan akan mengubah peradaban manusia",
		Title:             "The AI Revolution",
		Quantity:          14,
		AvailableQuantity: 14,
		PublishDate:       time.Now(),
		Edges: ent.BooksEdges{
			Categories: []*ent.Categories{
				{ID: 4}, // Technology
				{ID: 3}, // Science
				{ID: 2}, // Non-Fiction
			},
		},
	},
}

func BookSeeder(client *ent.Client) {
	var bookscreate []*ent.BooksCreate

	for _, book := range books {
		b := book

		create := client.Books.Create().
			SetAuthor(b.Author).
			SetTitle(b.Title).
			SetDescription(b.Description).
			SetQuantity(b.Quantity).
			SetAvailableQuantity(b.AvailableQuantity).
			SetPublishDate(b.PublishDate)

		for _, cat := range b.Edges.Categories {
			create = create.AddCategoryIDs(cat.ID)
		}

		bookscreate = append(bookscreate, create)
	}

	_, err := client.Books.CreateBulk(bookscreate...).Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
