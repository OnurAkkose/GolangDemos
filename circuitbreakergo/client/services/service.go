package services

import (
	"client/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiService struct {
	host    string
	port    int
	version int
}

func NewApiService(apihost string, apiPort int, apiVersion int) *ApiService {
	return &ApiService{
		host:    apihost,
		port:    apiPort,
		version: apiVersion,
	}
}

func (service *ApiService) SendRequest(c echo.Context, urlPath string) (interface{}, http.Header, error) {
	return common.SendRequest(service.host, service.port, urlPath, c)
}
