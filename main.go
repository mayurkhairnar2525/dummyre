package main

import "fmt"

func helloMsg(msg string) {
	fmt.Println(msg)
}

func sum(a, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

func hello() {
	fmt.Println("hello")
}

func main() {
	helloMsg("hello")
	sum(10, 15)
	hello()
}
