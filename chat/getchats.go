package chat

import "github.com/jinzhu/gorm"

//GetUserChats получает список чатов конкретного пользователя
func GetUserChats(db *gorm.DB, r *GetUserChatsRequest) (chats []*Chat, err error) {

	rows, err := db.Raw("SELECT id, name, created_at, users FROM chats WHERE  $1 = ANY (chats.users)", r.User).Rows()

	for rows.Next() {
		c := new(Chat)
		rows.Scan(&c.Id, &c.Name, &c.Created_at, &c.Users)
		chats = append(chats, c)
	}
	return chats, err
}
