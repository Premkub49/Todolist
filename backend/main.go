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

type Task struct {
	ID       int    `json:"id"`
	Listname string `json:"listname"`
	Deadline string `json:"deadline"`
	Detail   string `json:"detail"`
	Username string `json:"username"`
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
		"CREATE TABLE IF NOT EXISTS userdata(username text PRIMARY KEY NOT NULL,email text NOT NULL,password text NOT NULL);",
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS userlist(id SERIAL PRIMARY KEY NOT NULL,listname text NOT NULL,deadline timestamp with time zone NOT NULL,detail text ,username text NOT NULL, FOREIGN KEY (username) REFERENCES userdata(username) ON DELETE CASCADE);",
	)
	if err != nil {
		log.Fatal(err)
	}
	path := os.Getenv("PATH_TO_FRONT")
	app := fiber.New()
	log.Println(path)
	app.Use(cors.New(cors.Config{
		AllowOrigins: path,
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type",
	}))
	api := app.Group("/api")
	api.Post("/register", registerData)
	api.Post("/login", loginData)
	api.Post("/cookie", checkCookie)
	api.Post("/createtask", createTaskAPI)
	api.Post("/getUserTask", getUserTaskAPI)
	api.Delete("/deleteTask", deleteTaskAPI)
	api.Put("/updateTask", editTaskAPI)
	app.Listen(":8080")
}
