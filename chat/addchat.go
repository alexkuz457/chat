package chat

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//структура запроса на добавление пользователя
type Request struct {
	Name  string `json:"name"`
	Users []int  `json:"users"`
}

//структура чата в бд.
//Users      []user.User `gorm:"many2many:users;"` //users - список пользователей в чате, отношение многие-ко-многим
type Chat struct {
	Id         string    //id - уникальный идентификатор чата
	Name       string    //name - уникальное имя чата
	Users      []int     `gorm:"-"` //users - список пользователей в чате, отношение многие-ко-многим
	Created_at time.Time //created_at - время создания
}

//создает нового пользователя в базе данных
func CreateChat(db *gorm.DB, r *Request) (c *Chat, err error) {
	c = new(Chat)
	c.Name = r.Name
	c.Users = r.Users

	db.Create(c)
	db.Model(c).Exec("UPDATE public.chats SET users=$1 WHERE id=$2", pq.Array(c.Users), c.Id)

	return c, err
}