package main

import "fmt"

// Sorting algorithm

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		midIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[midIndex] {
				arr[j], arr[midIndex] = arr[midIndex], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 1, 3, 2, 0, 4}
	bubbleSort(arr)
	fmt.Println(arr)

	insertionSort(arr)
	fmt.Println(arr)

	selectionSort(arr)
	fmt.Println(arr)
}
