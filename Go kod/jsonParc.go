package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
	Job  string
	Info Info
}
type Info struct {
	Email      string
	Phone      string
	TelegramID int
}

func main() {
	// получили джейсон откуда-то
	jsonString := `[{"Name":"Jonh","Age":25,"Job":"fullstack","Info":{"Email":"lol@mail.com","Phone":"1231231","TelegramID":123123}},{"Name":"Sam","Age":27,"Job":"Frontend","Info":{"Email":"lol@mail.com","Phone":"1231231555","TelegramID":125553123}},{"Name":"Kevin","Age":28,"Job":"becend","Info":{"Email":"lol@mail.com","Phone":"123123231","TelegramID":1255513123}}]`
	// преобразуем строку в срез байт
	jsonData := []byte(jsonString)

	// создаем обьект класса юзер в который будем засовывать данные из джейсон
	users := []User{}

	// нужно перекодировать через АNMaRшала из джейсона в сруктуры
	json.Unmarshal(jsonData, &users)

	fmt.Println(users)

}
