package database

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type OutfitDB struct {
	Session *r.Session
}

//func (odb *OutfitDB) GetOutfit(s <-chan string) chan<- lib.Outfit {
//	id := <- s
//
//}
