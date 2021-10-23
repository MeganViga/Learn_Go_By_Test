package main
import "fmt"
const C = "Hello"
func Hello(name string)string{
	if name == ""{
		name = "World"
	}
	return C + ","+ name
}
func main(){
	fmt.Println(Hello("World"))
}