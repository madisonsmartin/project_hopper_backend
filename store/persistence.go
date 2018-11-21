package store

import (
	"database/sql"
	"log"

	"github.com/zachlefevre/project_hopper_backend/com"
)

type persistence interface {
	add(event *pb.Event) error
	getAll() []*pb.Event
	init()
}

type postgresDB struct {
}

func (db postgresDB) add(event *pb.Event) error {
	log.Printf("Adding event to DB", event)
	return nil
}

func (db postgresDB) getAll() []*pb.Event {
	log.Printf("request for all events")
	return nil
}

func (db postgresDB) init() {
	log.Printf("initializing DB")
	con, err := sql.Open("postgres", "postgresql://root@event-db:26257/?sslmode=disable")
	defer con.Close()
	if err != nil {
		log.Fatal("error connecting to the command DB: ", err)
	}

	if _, err := con.Exec(
		"CREATE DATABASE IF NOT EXISTS log"); err != nil {
		log.Fatal("cannot create database: ", err)
	}
	if _, err := con.Exec(
		"CREATE TABLE IF NOT EXISTS log.commands (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal("cannot create table: ", err)
	}
}
