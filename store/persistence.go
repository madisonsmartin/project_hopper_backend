package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/zachlefevre/project_hopper_backend/com"
)

const (
	connectionstring = "postgresql://root@cockroachdb-public:26257/?sslmode=disable"
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
	con, err := sql.Open("postgres", connectionstring)
	defer con.Close()
	if err != nil {
		log.Fatal("Failed to open DB connection")
	}
	eventString := fmt.Sprintf("%v, %v, %v, %v, %v, %v",
		event.EventId,
		event.EventType,
		event.AggregateId,
		event.AggregateType,
		event.EventData,
		event.Channel)
	if res, err := con.Exec("INSERT INTO log.commands VALUES(" + eventString + ")"); err != nil {
		log.Printf("Failed to persist to DB", err)
	} else {
		log.Printf("persisted event to DB", res)
	}
	return nil
}

func (db postgresDB) getAll() []*pb.Event {
	log.Printf("request for all events")
	return nil
}

func (db postgresDB) init() {
	log.Printf("initializing DB")
	con, err := sql.Open("postgres", connectionstring)
	defer con.Close()
	if err != nil {
		log.Fatal("error connecting to the command DB: ", err)
	}
	log.Printf("Checking if log database exists")
	if res, err := con.Exec(
		"CREATE DATABASE IF NOT EXISTS log"); err != nil {
		log.Fatal("cannot create database: ", err)
	} else {
		log.Println("created database", res)
	}
	log.Printf("Checking if commands db exists")
	if res, err := con.Exec(
		"CREATE TABLE IF NOT EXISTS log.commands (id INT PRIMARY KEY, string event_type, string aggregate_id, string aggregate_type, string event_data, string channel)"); err != nil {
		log.Fatal("cannot create table: ", err)
	} else {
		log.Printf("created table", res)
	}
}
