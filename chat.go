package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	User struct {
		Name string `json:"name"`
	}
)

func main() {
	e := echo.New()
	e.POST("users/", addUser)
	e.Logger.Fatal(e.Start(":9000"))
}

func addUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}
