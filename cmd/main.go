package main

import (
	handlers "finance/internal/spends/handlers"
	spendsService "finance/internal/spends/service"

	"github.com/labstack/echo/v4"
)

// crud для spend
// GIT init

// подключить postgress
// сделать htmx

func main() {

	repository := spendsService.NewSpendsRepository()
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
