package main

import (
	"encoding/json"
	"fmt"
)

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
