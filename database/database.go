package database

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Session  string
	Outfit   *OutfitDB
}

func (db *Database) load() {
	outfit := OutfitDB{}
	db.Outfit = &outfit
}

func (db *Database) OpenConnection() *Database {
	return db
}

func (db *Database) CreateTables() {

}
