package chat

import (
	"time"
)

//структура запроса на добавление пользователя
type Request struct {
	Name  string `json:"name"`
	Users []int  `json:"users"`
}

//структура запроса на получение чатов конкретного пользователя
type GetUserChatsRequest struct {
	User int `json:"user"`
}

//структура чата в бд.
//Users      []user.User `gorm:"many2many:users;"` //users - список пользователей в чате, отношение многие-ко-многим
type Chat struct {
	Id         string    //id - уникальный идентификатор чата
	Name       string    //name - уникальное имя чата
	Users      []int     `gorm:"-"` //users - список пользователей в чате, отношение многие-ко-многим
	Created_at time.Time //created_at - время создания
}
