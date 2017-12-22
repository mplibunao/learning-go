package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
