package main

import "fmt"

// sorting algo

func bubbleSort(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return 0
}

func insertionSort(arr []int) int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return 0
}

func selectionSort(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		midIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[midIndex] {
				arr[j], arr[midIndex] = arr[midIndex], arr[j]
			}
		}
	}
	return 0
}

func hello(msg string) {
	fmt.Println("Message:", msg)
}

func main() {
	arr := []int{50, 20, 40, 10, 30}
	bubbleSort(arr)
	fmt.Println(arr)

	insertionSort(arr)
	fmt.Println(arr)

	selectionSort(arr)
	fmt.Println(arr)

	hello("Hey!")
}
