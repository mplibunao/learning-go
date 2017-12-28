package main

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
}
