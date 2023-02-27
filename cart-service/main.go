package main

import (
	"cart/pkg/api"
	"cart/pkg/db"
	"cart/pkg/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// step 1: read config
	config, err := util.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	// step 2: connect to database
	dbName := config.DB_NAME
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DB_USERNAME, config.DB_PASSWORD, config.DB_URL, "5432", config.DB_DATABASE)
	conn, err := sql.Open(dbName, dbSource)

	if err != nil {
		log.Fatal(err)
	}

	// step 3: registe all db operations
	store := db.NewStore(conn)

	// step 4: Listen and Serve
	server := api.NewServer(config, store)
	err = server.Start(config.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
