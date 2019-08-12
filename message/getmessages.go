package message

import (
	"github.com/jinzhu/gorm"
)

//GetChatMessages получает список сообщений в чате
func GetChatMessages(db *gorm.DB, r *GetChatMessagesRequest) (messages []*Message, err error) {

	rows, err := db.Raw("SELECT id, chat, created_at, autor, text  FROM messages WHERE chat = $1)", r.Chat).Rows()

	for rows.Next() {
		m := new(Message)
		rows.Scan(&m.Id, &m.Chat, &m.Created_at, &m.Autor, &m.Text)
		messages = append(messages, m)
	}
	return messages, err
}
