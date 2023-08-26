package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var paymentPromise = make(map[string]interface{})

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(checkTokenMiddleware)
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://red-planet.com.co/", "138.128.160.43"},
	// 	AllowMethods: []string{http.MethodPost},
	// }))

	e.POST("/api/promise", createPromise)

	e.Start(":9000")
}

func createPromise(c echo.Context) error {
	if err := c.Bind(&paymentPromise); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	createPaymentPromise(paymentPromise)

	return c.String(http.StatusOK, "Promise Recieved")
}
