package server

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func initDatabase(dbDriver string, dbConnection string) (*sql.DB, error) {
	/* host := config.GetString("DB_HOST")
	port := config.GetString("DB_PORT")
	user := config.GetString("DB_USER")
	password := config.GetString("DB_PASSWORD")
	dbName := config.GetString("DB_NAME")
	dbDriver := config.GetString("DB_DRIVER")
	sslMode := config.GetString("SSL_MODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode) */

	db, err := sql.Open(dbDriver, dbConnection)
	if err != nil {
		log.Printf("Error while initializing database %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error while validation database connection: %v", err)
		return nil, err
	}

	return db, nil
}

func ConnectToDb() *sql.DB {
	var count int
	dbConnection := os.Getenv("DBCONNECTION")
	dbDriver := os.Getenv("DB_DRIVER")

	for {
		dbHandler, err := initDatabase(dbDriver, dbConnection)

		if err == nil {
			return dbHandler
		}

		log.Println("Postgres container not yet ready...")
		count++
		log.Println(count)

		if count > 10 {
			log.Fatalf("Error while initializing database %v", err)
			return nil
		}

		log.Println("Backing off for five seconds...")
		time.Sleep(5 * time.Second)
	}
}
