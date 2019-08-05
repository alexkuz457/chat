package message

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Структура сообщения в чате.
type Message struct {
	id         string    //id - уникальный идентификатор сообщения
	chat       string    //chat - ссылка на идентификатор чата, в который было отправлено сообщение
	autor      string    //author - ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
	text       string    //text - текст отправленного сообщения
	created_at time.Time //created_at - время создания
}

//структура запроса на добавление сообщения
type Request struct {
	chat  string `json:"chat"`
	autor string `json:"autor"`
	text  string `json:"text"`
}

//создает нового пользователя в базе данных
func CreateMessage(db *gorm.DB, r *Request) (m *Message, err error) {
	m = new(Message)

	db.Create(m)
	//db.Model(c).Exec("UPDATE public.chats SET users=$1 WHERE id=$2", pq.Array(c.Users), c.Id)

	return m, err
}
