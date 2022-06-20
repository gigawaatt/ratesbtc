package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "172.27.80.1"
	port     = 5432
	user     = "postgres"
	password = "user"
	dbname   = "postgres"
)

func Init() *sql.DB {

	// cfg := mysql.Config{
	// 	User:                 env.GetVars().GetDbUsername(),
	// 	Passwd:               env.GetVars().GetDbPasswd(),
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.0.1:3306",
	// 	DBName:               "rate",
	// 	AllowNativePasswords: true,
	// }

	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	psqlconn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user, password, host, port, dbname)

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
