package main

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	r "github.com/edwinyoyada/bopay/repositories"
	"github.com/edwinyoyada/bopay/services"
	"github.com/edwinyoyada/bopay/controllers"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DB_URI")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	VARepo := r.NewVARepo(db)
	VAService := services.NewVAService(VARepo)
	VAController := controllers.NewVAController(VAService)

	e.POST("/callbacks/virtual-accounts", VAController.UpdateVACallback)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}