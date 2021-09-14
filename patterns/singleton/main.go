package main

import "epam-mentoring/oop/patterns/singleton/net"

type Database struct {
	tcp *net.TCP
}

func NewDatabase(tcp *net.TCP) *Database {
	return &Database{tcp: tcp}
}

func (db *Database) Select() {
	db.tcp.GetConnection().Print()
}

func main() {
	tcp := net.NewTCP("шлях в епам")
	db := NewDatabase(tcp)
	db.Select()
	db.Select()
	db.Select()
	db.Select()
	db.Select()
	db.Select()
	db.Select()
}
