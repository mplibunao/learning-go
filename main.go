package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetAstronauts(g GetWebRequest) int {
	url := "http://api.open-notify.org/astros.json"
	bodyBytes := g.FetchBytes(url)
	peopleResult := people{}
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return peopleResult.Number
}

func main() {
	liveClient := LiveGetWebRequest{}
	number := GetAstronauts(liveClient)

	fmt.Println(number)
}
