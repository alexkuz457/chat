package message

import (
	"time"
)

//Message Структура сообщения в чате.
type Message struct {
	Id         int       //id - уникальный идентификатор сообщения
	Chat       int       //chat - ссылка на идентификатор чата, в который было отправлено сообщение
	Autor      int       //author - ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
	Text       string    //text - текст отправленного сообщения
	Created_at time.Time //created_at - время создания
}

//Request структура запроса на добавление сообщения
type Request struct {
	Chat  int    `json:"chat"`
	Autor int    `json:"author"`
	Text  string `json:"text"`
}

//GetUserMessagesRequest структура запроса на получение всех сообщений пользователя
type GetChatMessagesRequest struct {
	Chat int `json:"chat"`
}
