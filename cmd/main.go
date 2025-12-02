package main

import (
	db "finance/internal/db"
	handlers "finance/internal/spends/handlers"
	spendsService "finance/internal/spends/service"
	"log"

	"github.com/labstack/echo/v4"
)

// подключить postgress
// Добавить мииграции

// сделать HTMX

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DataBase: %v", err)
	}

	repository := spendsService.NewSpendsRepository(database)
	service := spendsService.NewSpendsService(repository)

	Handlers := handlers.NewSpendsHandlers(service)

	e := echo.New()
	e.GET("/spends", Handlers.GetAllSpends)
	e.GET("/spends/:id", Handlers.GetSpend)
	e.POST("/spends", Handlers.PostSpend)
	e.PATCH("/spends/:id", Handlers.PatchSpend)
	e.DELETE("/spends/:id", Handlers.DeleteSpend)

	e.Logger.Fatal(e.Start(":8080"))

}
