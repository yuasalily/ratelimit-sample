package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID int `json:"id"`
}

func main() {
	e := echo.New()

	e.GET("/", noop)
	e.GET("/user", getUser)
	e.Start(":8080")
}

func noop(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, &User{ID: 1})
}
