package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker // 1
func init() {
	var settings gobreaker.Settings // 2
	settings.Name = "HTTP GET"
	settings.ReadyToTrip = func(counts gobreaker.Counts) bool {
		// circuit breaker will trip when 60% of requests failed and at least 3 requests were made
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	settings.Timeout = time.Second * 10
	settings.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		if to == gobreaker.StateOpen {
			fmt.Println("State Open!")
		}
		if from == gobreaker.StateOpen && to == gobreaker.StateHalfOpen {
			fmt.Println("Going from Open to Half Open")
		}
		if from == gobreaker.StateHalfOpen && to == gobreaker.StateClosed {
			fmt.Println("Going from Half Open to Closed!")
		}
	}
	cb = gobreaker.NewCircuitBreaker(settings)
}

func SendRequest(host string, port int, urlPath string, c echo.Context) (interface{}, http.Header, error) {

	url := fmt.Sprintf("%s://%s:%d%s", "http", host, port, urlPath)

	var response interface{}

	u := new(interface{})
	if err := c.Bind(u); err != nil {

		return response, nil, err
	}
	requestBody, err := json.Marshal(u)
	if err != nil {
		return response, nil, err
	}
	/////
	body, err := cb.Execute(func() (interface{}, error) {
		res, err := http.Post(url, "application/json; charset=UTF-8", bytes.NewBuffer(requestBody))
		defer func() {
			if res != nil {
				res.Body.Close()
			}
		}()
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		if len(data) > 0 {
			if err := json.Unmarshal(data, &response); err != nil {
				return nil, err
			}
		}
		if res.StatusCode != http.StatusOK {
			return nil, err
		}
		return response, nil
		////
	})

	if err != nil {
		return response, nil, err
	}
	return body.([]byte), c.Response().Header(), nil
}
