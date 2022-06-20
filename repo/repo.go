package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {

	host := GetVars().GetDbUri()
	user := GetVars().GetDbUsername()
	password := GetVars().GetDbPasswd()
	dbname := GetVars().GetDBaddr()

	psqlconn := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
		user, password, host, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(err)
	}

	log.Println("Connection to database established")

	return db
}
