package main

import (
	"net/http"
	"os/user"
	"time"

	//"github.com/alexkuz457/chat/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
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

//Сообщение в чате. Имеет следующие свойства:
//id - уникальный идентификатор сообщения
//chat - ссылка на идентификатор чата, в который было отправлено сообщение
//author - ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
//text - текст отправленного сообщения
//created_at - время создания
type Message struct {
	id         string `json:"id"`
	chat       string `json:"chat"`
	autor      string `json:"autor"`
	text       string `json:"text"`
	created_at string `json:"created_at"`
}

func main() {
	e := echo.New()
	e.POST("users/add", addUser)
	e.POST("chats/add", addChat)
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
	//user.AddUser(c)

	u := new(user.Request)

	if err = c.Bind(u); err != nil {
		return
	}
	//db := Database()
	//db.Save(u)
	return c.JSON(http.StatusOK, "u")
}

func addChat(c echo.Context) (err error) {

	u := new(Chat)
	db := Database()

	if err = c.Bind(u); err != nil {
		return
	}
	db.Create(u)
	//query := "INSERT INTO public.chats(name, users) VALUES ($1, $2)"
	//_, err = db.Exec(query, u.Name, pq.Array(u.Users))
	//db.Save(u)

	return c.JSON(http.StatusOK, u.Id)
}
