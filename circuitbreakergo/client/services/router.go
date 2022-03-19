package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (service *ApiService) GetAllUserDefinition(c echo.Context) error {
	response, _, err := service.SendRequest(c, "/users/get")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK, response)
}
