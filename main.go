package main

import "fmt"

func git() {
	fmt.Println("Hello, From git")
}

func sayHello(msg string) {
	fmt.Println(msg)
}

func main() {
	git()
	sayHello("hi")
}
