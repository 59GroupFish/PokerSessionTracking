package main

import (
	"log"
	"net/http"
	"os"

	"github.com/59GroupFish/PokerSessionTracking/db"
	"github.com/59GroupFish/PokerSessionTracking/rest"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
)

func main() {
	SetEnvVars()

	if err := db.InitDB(); err != nil {
		log.Fatalf("FATAL: Failed to initialize database: %v", err)
	}
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Printf("WARNING: Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}()

	e := echo.New()
	err := Handler(e)
	if err != nil {
		log.Fatalf("Error setting up routes: %v", err)
	}

	s := http.Server{
		Addr:    ":8081",
		Handler: e,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func SetEnvVars() {
	os.Setenv("GOOSE_DBSTRING", "host=localhost user=postgres password=ah89213ohr0 dbname=poker_session_tracker")
	os.Setenv("GOOSE_DRIVER", "pgx")
	os.Setenv("GOOSE_DIR", "db/migrations")
	os.Setenv("DATABASE_CONN", "host=localhost user=postgres password=ah89213ohr0 dbname=poker_session_tracker")
}

func Handler(e *echo.Echo) error {
	// add auth endpoint

	rest.FrontEndRoutes(e)
	rest.BackendRoutes(e)

	return nil
}
