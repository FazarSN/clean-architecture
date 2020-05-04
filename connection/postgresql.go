package connection

import (
	"database/sql"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewPostgreConn() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL") //postgres://{username}:{password}@{host}:{port}/{dbName}
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	return dbConn
}
