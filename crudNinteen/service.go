package crudninteen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Sample() {
	fmt.Println("Learning CRUD...")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatal("Error making GET request:", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while getting response", err)
		return
	}
	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error while reading response", err)
	// 	return
	// }
	// fmt.Println(string(data)) // but this is in json format, and if you want to use it then you have to unmarhsall it

	// 2nd option
	var todos Todo
	errs := json.NewDecoder(res.Body).Decode(&todos)
	if errs != nil {
		log.Fatal("Error while decoding response", err)
	}
	log.Fatal(todos)

}
