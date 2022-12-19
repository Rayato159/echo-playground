package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Rayato159/go-echo/modules/books/handlers"
	"github.com/Rayato159/go-echo/modules/books/repositories"
	"github.com/Rayato159/go-echo/modules/books/usecases"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Config struct {
	App App
}

type App struct {
	Host string
	Port string
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error, can't load .env with an error: %s", err.Error())
	}

	cfg := Config{
		App: App{
			Host: os.Getenv("ECHO_HOST"),
			Port: os.Getenv("ECHO_PORT"),
		},
	}
	appUrl := fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)

	e := echo.New()

	// Modules
	v := "v1"

	// *Books
	booksRep := repositories.NewBooksRepository()
	booksUse := usecases.NewBooksUsecase(booksRep)
	handlers.NewBooksHandler(e, v, booksUse)

	e.Logger.Fatal(e.Start(appUrl))
}
