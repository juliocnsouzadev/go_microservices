package main

import (
	"authetication/cmd/api/config"
	"authetication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	conn := connectToDB()
	if conn == nil {
		log.Panic("Cannot connect to Postgres")
	}

	app := config.Config{
		DB:     conn,
		Models: data.New(conn),
	}

	log.Printf("Starting auth service on port %s\n", webPort)

	//define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress not yet ready")
			counts++
		} else {
			log.Println("Connected to Postgress")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Waiting for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}

}
