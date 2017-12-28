package main

import (
	"encoding/json"
	"fmt"
	"log"
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

type GetWebRequest interface {
	FetchBytes(url string) []byte
}

func GetAstronauts(getWebRequest GetWebRequest) int {
	url := "http://api.open-notify.org/astros.json"
	bodyBytes := getWebRequest.FetchBytes(url)
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

	// spaceClient := http.Client{
	// 	Timeout: time.Second * 2, // Max of 2 secs
	// }

	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req.Header.Set("User-Agent", "spacecount")

	// res, getErr := spaceClient.Do(req)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }

	// fmt.Println("HTTP: %s\n", res.Status)

	// body, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }

	// people1 := people{}
	// jsonErr := json.Unmarshal(body, &people1)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }

	// fmt.Println(people1.Number)
}
