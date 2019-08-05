package user

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//запрос на добавление пользователя
type Request struct {
	Username string `json:"username"`
}

//ответ при успешном добавлении пользователя
type Response struct {
	User *User
}

//User : структура содержащая данныe пользователя
type User struct {
	//gorm.Model
	Id         int16     `gorm:"primary_key" json:"id"`
	Username   string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

//создает нового пользователя в базе данных
func CreateUser(db *gorm.DB, r *Request) (u *User, err error) {
	u = new(User)
	u.Username = r.Username
	//здесь должны быть проверки
	db.Save(u)
	return u, err
}
