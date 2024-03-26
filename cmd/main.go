package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamideh/go_take_exam/routes"
	_ "github.com/lib/pq"

	// "gorm.io/driver/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "gorm.io/gorm"
)

// Define a struct to represent your database model
type User struct {
	ID   uint
	Name string
	Age  uint
}

func main() {
	// Define the database configuration

	// dsn := "postgres://hamiltoon:1234@localhost:5432/examtaker_db"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Failed to connect to database:", err)
	// 	return
	// }

	// dsn2 := url.URL{
	// 	User:     url.UserPassword("hamiltoon", "1234"),
	// 	Scheme:   "postgres",
	// 	Host:     fmt.Sprintf("%s:%d", "localhost", "5432"),
	// 	Path:     "examtaker_db",
	// 	RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	// }
	// db, err := gorm.Open("postgres", dsn2.String())
	// // Ping the database to ensure the connection is established
	// if sqlDB, err := db.DB(); err != nil {
	// 	fmt.Println("Failed to get database instance:", err)
	// 	return
	// } else if err := sqlDB.Ping(); err != nil {
	// 	fmt.Println("Failed to ping database:", err)
	// 	return
	// }

	// fmt.Println("Database connection established successfully")

	// Migrate the schema (automatically create the table based on the model)
	// db.AutoMigrate(&User{})

	// Perform database operations (e.g., CRUD operations)
	// For example:
	// var users []User
	// db.Find(&users)

	// Close the underlying database connection when you're done
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	fmt.Println("Failed to get database instance:", err)
	// 	return
	// }
	// sqlDB.Close()
	r := mux.NewRouter()
	routes.RegisterExamTaker(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

// type application struct {
// 	appName string
// 	server  server
// 	debug   bool
// 	errLog  *log.Logger
// 	infoLog *log.Logger
// 	view    *jet.set
// 	Models  models.Models
// }

// type server struct {
// 	host string
// 	port string
// 	url  string
// }

// func main() {
// 	// Replace the connection details (user, dbname, password, host) with your own
// 	db, err := sqlx.Connect("postgres", "user=hamiltoon dbname=examtaker_db sslmode=disable password=1234 host=localhost")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer db.Close()

// 	// Test the connection to the database
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println("Successfully Connected to Postgresql___ examtaker db__ hamiltoon user")
// 	}
// }
