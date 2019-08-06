package chat

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//получает список чатов конкретногопользователя
func GetUserChats(db *gorm.DB, r *Request) (c *Chat, err error) {
	c = new(Chat)
	c.Name = r.Name
	c.Users = r.Users

	db.Create(c)
	db.Model(c).Exec("UPDATE public.chats SET users=$1 WHERE id=$2", pq.Array(c.Users), c.Id)

	return c, err
}
