package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//CreateUser создает нового пользователя в базе данных
func CreateUser(db *gorm.DB, r *Request) (u *User, err error) {
	u = new(User)
	u.Username = r.Username
	//здесь должны быть проверки
	db.Create(u)
	return u, err
}
