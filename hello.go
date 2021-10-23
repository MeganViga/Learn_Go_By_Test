package main

//import "fmt"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix  = "Bonjour"
func Hello(name string, language string)string{
	if name == ""{
		name = "World"
	}
	
	return greetings(language) + ","+ name
}
func greetings(language string)(prefix string){
	switch language{
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

