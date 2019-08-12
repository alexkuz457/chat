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

//Database инициализирует и возвращает соединение с postgres
func Database() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=chat password=qwerty12 sslmode=disable")
	if err != nil {
		println("Ошибка подключения к базе данных %s", err)
	}
	return db, err
}

func addUser(c echo.Context) (err error) {
	r := new(user.Request)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, r)
	}
	db, err := Database()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	u, err := user.CreateUser(db, r)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer db.Close()

	return c.JSON(http.StatusOK, u.Id)
}

func addChat(c echo.Context) (err error) {
	r := new(chat.Request)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db, err := Database()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	u, err := chat.CreateChat(db, r)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, u.Id)
}

func addMessage(c echo.Context) (err error) {
	r := new(message.Request)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db, err := Database()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := message.CreateMessage(db, r)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, id)
}

func getChat(c echo.Context) (err error) {
	r := new(chat.GetUserChatsRequest)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db, err := Database()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	u, err := chat.GetUserChats(db, r)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer db.Close()

	return c.JSON(http.StatusOK, u)
}

func getMessages(c echo.Context) (err error) {
	r := new(message.GetChatMessagesRequest)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db, err := Database()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	u, err := message.GetChatMessages(db, r)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer db.Close()

	return c.JSON(http.StatusOK, u)
}
