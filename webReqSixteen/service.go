package webreqsixteen

import (
	"fmt"
	"io"
	"net/http"
)

func Sample() {
	fmt.Println("Hello world form web")
	httpRes, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting get res", err)
		return
	}
	defer httpRes.Body.Close()

	fmt.Printf("oo :%T\n", httpRes) // it give will give "*http.Response"
	data, err := io.ReadAll(httpRes.Body)
	if err != nil {
		fmt.Println("Error reading res", err)
		return
	}
	fmt.Println(string(data))
}
