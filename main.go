package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// Since body has Read and Close interface, it means it also has access to Read method
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// os.Stdout is type File which implements the writer interface
	io.Copy(lw, resp.Body)
}

// logWriter now implements the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes:", len(bs))
	return len(bs), nil
}
