package message

import (
	"github.com/jinzhu/gorm"
)

//CreateMessage создает нового пользователя в базе данных
func CreateMessage(db *gorm.DB, r *Request) (id int, err error) {
	m := new(Message)
	m.Autor = r.Autor
	m.Chat = r.Chat
	m.Text = r.Text
	db.Create(m)
	return m.Id, err
}
