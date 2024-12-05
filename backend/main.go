package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	host         = "postgres"
	port         = 5432
	databaseName = os.Getenv("POSTGRES_DB")
	username     = os.Getenv("POSTGRES_USER")
	password     = os.Getenv("POSTGRES_PASSWORD")
)

var db *sql.DB
var err error

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)
	db, err = sql.Open("postgres", psqlInfo)
	_ = db
	_ = err
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Success")
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS userdata(username text PRIMARY KEY NOT NULL,email text NOT NULL,password text NOT NULL)",
	)
	if err != nil {
		log.Fatal(err)
	}
	path := os.Getenv("PATH")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: path,
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type",
	}))
	api := app.Group("/api")
	api.Post("/register", registerData)
	api.Post("/login", loginData)
	app.Listen(":8080")
}
