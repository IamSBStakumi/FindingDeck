package main

import (
	"net/http"

	project "github.com/IamSBStakumi/findingdeck/internal/modules/project/interface"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	projectHandler := project.NewHTTPHandler()
	projectHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}