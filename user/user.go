package user

import (
	"net/http"
	"os/user"

	"github.com/labstack/echo"
)

//запрос на добавление пользователя
type Request struct {
	Username string `json:"name"`
}

//ответ при успешном добавлении пользователя
type Response struct {
	User *user.User
}

//создает нового пользователя в базе данных
func CreateUser(c echo.Context) (err error) {
	u := new(Request)
	if err = c.Bind(u); err != nil {
		return
	}
	//db := Database()
	//db.Save(u)
	return c.JSON(http.StatusOK, u)
}
