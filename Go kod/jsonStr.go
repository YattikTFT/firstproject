package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"nameOFDream"`
	Age  int    `json:"ageOFDream"`
	Job  string `json:"JObOFDream"`
	Info Info   `json:"InformationOFDream"`
}
type Info struct {
	Email      string
	Phone      string
	TelegramID int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		user1 := User{Name: "Jonh", Age: 25, Job: "fullstack", Info: Info{Email: "lol@mail.com", Phone: "1231231", TelegramID: 123123}}
		user2 := User{Name: "Sam", Age: 27, Job: "Frontend", Info: Info{Email: "lol@mail.com", Phone: "1231231555", TelegramID: 125553123}}
		user3 := User{Name: "Kevin", Age: 28, Job: "becend", Info: Info{Email: "lol@mail.com", Phone: "123123231", TelegramID: 1255513123}}

		//список ЮЗеров
		users := []User{}

		// добавляем юзеров в масив
		users = append(users, user1)
		users = append(users, user2)
		users = append(users, user3)

		// Переганяем данные в Джейсон
		jsonData, _ := json.Marshal(users)

		jsonString := string(jsonData)
		// байты в строки
		fmt.Println(jsonString)
		//выводим на экран
		fmt.Fprintf(w, jsonString)
	})
	http.ListenAndServe(":8080", nil)
}
