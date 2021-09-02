package links

import (
	"log"

	db "github.com/neurocome/ine_go/internal/pkg/db/pgsql"
	"github.com/neurocome/ine_go/internal/users"
)

type Link struct {
	ID      string     `pg:"id"`
	Title   string     `pg:"title"`
	Address string     `pg:"address"`
	UserID  int        `pg:"user_id"`
	User    users.User `pg:"rel:has-one"`
}

func (link *Link) Save() int64 {
	id, err := db.Db.Model(link).Insert()
	if err != nil {
		panic(err)
	}
	log.Print("Row Inserted")
	return int64(id.RowsReturned())
}

func GetAll() []Link {
	link := new([]Link)
	err := db.Db.Model(link).Relation("User").Select()
	if err != nil {
		panic(err)
	}
	return *link
}
