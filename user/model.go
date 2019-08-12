package user

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Request структура запроса на добавление пользователя
type Request struct {
	Username string `json:"username"`
}

//Response ответ при успешном добавлении пользователя
type Response struct {
	User *User
}

//User : структура содержащая данныe пользователя
type User struct {
	Id         int16     `gorm:"primary_key" json:"id"`
	Username   string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}
