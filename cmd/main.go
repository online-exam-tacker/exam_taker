// how to run : CGO_ENABLED=0 go run main.go
// how to run : CGO_ENABLED=0 go build main.go

package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// type application struct {
// 	appName string
// 	server  server
// 	debug   bool
// 	errLog  *log.Logger
// 	infoLog *log.Logger
// 	view    *jet.set
// 	Models  models.Models
// }

type server struct {
	host string
	port string
	url  string
}

func main() {
	// Replace the connection details (user, dbname, password, host) with your own
	db, err := sqlx.Connect("postgres", "user=hamiltoon dbname=examtaker_db sslmode=disable password=1234 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected to Postgresql___ examtaker db__ hamiltoon user")
	}
}
