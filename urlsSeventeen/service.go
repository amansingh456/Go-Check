package urlsseventeen

import (
	"fmt"
	"net/url"
)

func Sample() {
	fmt.Println("hello world, from urls")

	myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error while parsing eurl", err)
		return
	}
	fmt.Printf("type of url: %T\n", parsedUrl)
	fmt.Println(parsedUrl)
	fmt.Println("Scheme", parsedUrl.Scheme)
	fmt.Println("Host", parsedUrl.Host)
	fmt.Println("Path", parsedUrl.Path)
	fmt.Println("RawQuery", parsedUrl.RawQuery)

	// modifying url f

	parsedUrl.Host = "aman.com"
	// like this we can change the  everything in url

	fmt.Println(parsedUrl)

}
