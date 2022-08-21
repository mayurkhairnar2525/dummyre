package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}

func sayHello(msg string) {
	fmt.Println(msg)
}

func main() {
	hello()
	sayHello("hello world")
}
