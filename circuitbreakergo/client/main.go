package main

import (
	"client/services"

	"github.com/labstack/echo/v4"
)

func main() {
	apiService := services.NewApiService("localhost", 6161, 1)
	e := echo.New()
	e.POST("/api/v1/users/get", apiService.GetAllUserDefinition)
	e.Logger.Fatal(e.Start(":1323"))
}
