package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("What's your name?")
	fmt.Scan(&name)
	fmt.Println("Hello", name)

	var age uint
	fmt.Println("What's your age?")
	fmt.Scan(&age)
	fmt.Printf("Your name is %s and your age is %d.", name, age)

	newMessage := fmt.Sprintf("We hear that your name is %s and that you are %d years old.", name, age)
	fmt.Println(newMessage)

}
