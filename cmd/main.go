package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

// type server struct {
// 	host string
// 	port string
// 	url  string
// }

// Define a struct to represent your database model
type User struct {
	ID   uint
	Name string
	Age  uint
}

func main() {
	// Define the database configuration
	dsn := "user=hamiltoon password=1234 dbname=examtaker_db host=localhost port=5432 sslmode=disable TimeZone=UTC"

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Ping the database to ensure the connection is established
	if sqlDB, err := db.DB(); err != nil {
		fmt.Println("Failed to get database instance:", err)
		return
	} else if err := sqlDB.Ping(); err != nil {
		fmt.Println("Failed to ping database:", err)
		return
	}

	fmt.Println("Database connection established successfully")

	// Migrate the schema (automatically create the table based on the model)
	db.AutoMigrate(&User{})

	// Perform database operations (e.g., CRUD operations)
	// For example:
	// var users []User
	// db.Find(&users)

	// Close the underlying database connection when you're done
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get database instance:", err)
		return
	}
	sqlDB.Close()

}

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
