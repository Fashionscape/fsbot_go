package database

import (
	"github.com/salmonllama/fsbot_go/lib"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
)

type Database struct {
	Host string
	Port string
	Name string
	User string
	Password string
	Session *r.Session
	Outfit *OutfitDB
}

func (db *Database) load() {
	outfit := OutfitDB{Session:db.Session}
	db.Outfit = &outfit
}

func (db *Database) OpenConnection() *Database {
	session, err := r.Connect(r.ConnectOpts{
		Address:             db.Host + ":" + db.Port,
		Database:            db.Name,
		Username:            db.User,
		Password:            db.Password,
	})
	lib.Check(err)
	db.Session = session

	db.load()
	return db
}

func (db *Database) CreateTables() {
	var (
		result r.WriteResponse
		err error
	)

	result, err = r.TableCreate("outfits").RunWrite(db.Session)
	if err == nil {
		log.Print(result)
	}

	result, err = r.TableCreate("guild_conf").RunWrite(db.Session)
	if err ==  nil {
		log.Print(result)
	}

	result, err = r.TableCreate("guild_blacklist").RunWrite(db.Session)
	if err == nil {
		log.Print(result)
	}

	result, err = r.TableCreate("user_blacklist").RunWrite(db.Session)
	if err == nil {
		log.Print(result)
	}

	result, err = r.TableCreate("color_roles").RunWrite(db.Session)
	if err == nil {
		log.Print(result)
	}
}