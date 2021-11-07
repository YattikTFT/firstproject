package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Car struct {
	Mark  string `json:"Mark"`
	Model string `json:"Model"`
	Photo string `json:"Photo"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		markCar := r.URL.Query().Get("mark")
		modelCar := r.URL.Query().Get("model")
		photoCar := r.URL.Query().Get("photo")
		b := Car{Mark: markCar, Model: modelCar, Photo: photoCar}
		if markCar != "" {
			// сначала считываем то что была в файле
			dataFromFile, _ := ioutil.ReadFile("data.json")

			// создаем данные под джейон
			addCar := []Car{}

			json.Unmarshal(dataFromFile, &addCar)
			// ADD CAR  К СПИСКУ ТЕХ ЧТО БЫЛИ
			addCar = append(addCar, b)
			// собираем Json из обновленоо сиска
			jsonData, _ := json.Marshal(addCar)
			// с байтов в строку
			jsonstring := string(jsonData)

			// преоразовать json iz faile  v срез машин

			//открываем файл для работы с ним
			file, _ := os.Create("data.json")
			// отлоеное закрытие ФАЙЛА
			defer file.Close()

			//записываем в файл данные

			file.WriteString(jsonstring)
			// VOOD ON EKRAN

			fmt.Fprintf(w, jsonstring)

		}

	})
	http.ListenAndServe(":8080", nil)
}
