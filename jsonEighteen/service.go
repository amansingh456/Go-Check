package jsoneighteen

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsMarried bool   `json:"isMarried"`
}

func Sample() {
	fmt.Println("Hello from JSON")
	newPerson := Person{Name: "AmanSingh", Age: 24, IsMarried: false}
	fmt.Println(newPerson)

	// converting  newPerson into JSON

	marshalledData, err := json.Marshal(newPerson)
	if err != nil {
		fmt.Println("error while marshalling data", err)
		return
	}
	fmt.Println(string(marshalledData))

	//unmarshalling json data
	var newUser Person
	errs := json.Unmarshal(marshalledData, &newUser)
	if errs != nil {
		fmt.Println("error while unMarshalling data", errs)
		return
	}
	fmt.Println(newUser)

}
