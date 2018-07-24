package main

import (
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"strconv"
)

var (
	errorRate = 0.5
	randomz   = rand.New(rand.NewSource(42)) // ensure consistent
)

func fails() bool {
	// Asserting the source is uniformly distributed,
	// The random will then give numbers less than the expected rate in
	// "expected rate" percent of cases.
	return randomz.Float64() < errorRate
}

func handleReq() echo.HandlerFunc {
	return func(c echo.Context) error {
		if fails() {
			return c.JSON(http.StatusInternalServerError, "Error error")
		}
		return c.JSON(http.StatusOK, "world\n")
	}
}

func handleRateSet(c echo.Context) error {
	newRate := c.FormValue("rate")
	rate, err := strconv.ParseFloat(newRate, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a float")
	}
	if rate > 1.0 || rate < 0.0 {
		return c.JSON(http.StatusBadRequest, "Rate out of range")
	}

	errorRate = rate
	return c.JSON(http.StatusOK, errorRate)
}

func handleRateGet(c echo.Context) error {
	return c.JSON(http.StatusOK, errorRate)
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy\n")
}
func readyz(c echo.Context) error {
	return c.JSON(http.StatusOK, "ready\n")
}

func StartServer() {
	e := echo.New()

	e.Any("/", handleReq())
	e.POST("/errorrate", handleRateSet)
	e.GET("/errorrate", handleRateGet)
	e.GET("/healthz", healthz)
	e.GET("/readyz", readyz)

	e.Start(":9001")
}

func main() {
	StartServer()
}
