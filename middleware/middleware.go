package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func UseMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:800", "hieutran.io"},
	}))
}
