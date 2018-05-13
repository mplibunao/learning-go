package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type LiveGetWebRequest struct {
}

func (LiveGetWebRequest) FetchBytes(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Max 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
