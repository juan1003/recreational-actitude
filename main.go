package main

import "fmt"

func greet(name string) {
	greeting := fmt.Sprintf("Hello, %s", name)
	fmt.Println(greeting)	
}

func main() {
	greet("Juan De Jesus")
}
