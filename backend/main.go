package main

import (
	"net/http"
	"voucher-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})
	routes := routes.MakeRouter()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":5000"))
}
