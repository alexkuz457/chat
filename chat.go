package main

import (
	"net/http"

	"github.com/alexkuz457/chat/chat"
	"github.com/alexkuz457/chat/message"
	"github.com/alexkuz457/chat/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("users/add", addUser)
	e.POST("chats/add", addChat)
	e.POST("messages/add", addMessage)
	e.POST("chats/get", getChat)
	e.POST("messages/get", getMessages)
	e.Logger.Fatal(e.Start(":9000"))
}

func Database() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=chat password=qwerty12 sslmode=disable")
	if err != nil {
		println("Ошибка подключения к базе данных %s", err)
	}
	return db
}

func addUser(c echo.Context) (err error) {
	r := new(user.Request)
	if err = c.Bind(r); err != nil {
		return
	}
	db := Database()
	u, err := user.CreateUser(db, r)

	defer db.Close()

	return c.JSON(http.StatusOK, u.Id)
}

func addChat(c echo.Context) (err error) {
	r := new(chat.Request)
	if err = c.Bind(r); err != nil {
		return
	}
	db := Database()
	u, err := chat.CreateChat(db, r)
	defer db.Close()
	return c.JSON(http.StatusOK, u.Id)
}

func addMessage(c echo.Context) (err error) {
	r := new(message.Request)
	if err = c.Bind(r); err != nil {
		return
	}
	db := Database()
	id, err := message.CreateMessage(db, r)
	defer db.Close()
	return c.JSON(http.StatusOK, id)
}

func getChat(c echo.Context) (err error) {
	r := new(chat.GetUserChatsRequest)
	if err = c.Bind(r); err != nil {
		return
	}
	db := Database()
	u, _ := chat.GetUserChats(db, r)
	defer db.Close()

	return c.JSON(http.StatusOK, u)
}

func getMessages(c echo.Context) (err error) {
	r := new(message.GetChatMessagesRequest)
	if err = c.Bind(r); err != nil {
		return
	}
	db := Database()
	u, _ := message.GetChatMessages(db, r)
	defer db.Close()

	return c.JSON(http.StatusOK, u)
}
