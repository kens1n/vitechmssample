package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
	"log"
	"os"
)

const directory = "migrations"

func main() {
	path, _ := os.Getwd()
	godotenv.Load(path + "/.env")

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
