package main

import (
	"github.com/SamirMarin/user-management-service/internal/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/create", create)
	e.POST("/get", get)
	e.Logger.Fatal(e.Start(":1324"))
}

func create(c echo.Context) error {
	user := user.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := user.CreateUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "User created")
}

func get(c echo.Context) error {
	workout := user.User{}
	if err := c.Bind(&workout); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := workout.GetUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, workout)
}
