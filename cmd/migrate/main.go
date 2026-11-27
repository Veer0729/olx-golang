package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Veer0729/olx_golang/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up> | <down>")
	}

	cfg:= config.MustLoad()

	m, err := migrate.New(
        "file://migrations",
        cfg.DatabaseUrl, // reads .env for DB url
	)
	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		fmt.Println("up called")
		if err := m.Up(); err != nil {
			log.Fatalf("Error while up called: %v",err)
		}

	case "down":
		fmt.Println("down called")
		if err := m.Steps(-1); err != nil { // m.Down() would erase all our data and go to the default state: NOT RECEMONDED for production
			log.Fatalf("Error while down called: %v",err)
		}
	
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("running migration")
}

