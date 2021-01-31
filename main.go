package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/kafka_rest_handler_huge_amount_of_event/handler"
	mid "github.com/kafka_rest_handler_huge_amount_of_event/middleware"
)

func main() {
	e := echo.New();
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's ok bro !")
	})
	mid.UseMiddlewares(e)
	g := e.Group("/events")
	h := handler.New()
	h.RegisterEventHandler(g)

	e.Logger.Fatal(e.Start(":3000"))
}
