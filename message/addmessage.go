package message

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Структура сообщения в чате.
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

//CreateMessage создает нового пользователя в базе данных
func CreateMessage(db *gorm.DB, r *Request) (id int, err error) {
	m := new(Message)
	m.Autor = r.Autor
	m.Chat = r.Chat
	m.Text = r.Text
	db.Create(m)
	//db.Model(c).Exec("UPDATE public.chats SET users=$1 WHERE id=$2", pq.Array(c.Users), c.Id)

	return m.Id, err
}
