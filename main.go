package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func handleReq() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Its alive\n")
	}
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy\n")
}
func readyz(c echo.Context) error {
	return c.JSON(http.StatusOK, "ready\n")
}

func main() {
	e := echo.New()

	e.Any("/", handleReq())
	e.GET("/healthz", healthz)
	e.GET("/readyz", readyz)

	e.Start(":9001")
}
