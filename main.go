package main

import (
	"encoding/json"
	"fmt"
)

/*
In order to use a struct to unmarshal the JSON text,
we need to decorate it with tags that help the std library
understand how to map the properties

The property names need to begin with a capital which marks them as exportable or public

If Struct's property is the same in the JSON, you can skip the annotation
*/
type people struct {
	Number int `json:"number"`
}

func main() {
	text := `{
				"people": [
					{ "craft": "ISS", "name": "Sergey Rizhikov" },
					{ "craft": "ISS", "name": "Andrey Borisenk" }
				],
				"message": "success",
				"number": 6
			}`
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)
}
