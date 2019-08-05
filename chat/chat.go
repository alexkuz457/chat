package chat

import (
	"os/user"
	"time"
)

//Отдельный чат. Имеет следующие свойства:
//id - уникальный идентификатор чата
//name - уникальное имя чата
//users - список пользователей в чате, отношение многие-ко-многим
//created_at - время создания
type Chat struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Users      []user.User `gorm:"many2many:users;" json:"users"`
	Created_at time.Time   `json:"created_at"`
}
