package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := make([]int, 1000000)

	for i := 0; i < 999999; i++ {
		list[i] = rand.Intn(2000000)
	}

	start := time.Now()
	Quicksort(list)
	duration := time.Since(start)

	fmt.Println(duration)
}

func Quicksort(array []int) {
	size := len(array)

	if size <= 1 {
		return
	}

	left, right := 0, size-1
	middle := (left + right) / 2

	// Average of three
	if array[middle] < array[left] {
		swap(array, middle, left)
	}

	if array[right] < array[left] {
		swap(array, right, left)
	}

	if array[right] < array[middle] {
		swap(array, middle, right)
	}

	pivot := right - 1
	swap(array, middle, pivot)

	// Partition
	for i := 0; i < right; i++ {
		if array[i] < array[right] {
			swap(array, left, i)
			left++
		}
	}

	swap(array, right, left)

	Quicksort(array[:left])
	Quicksort(array[left+1:])
}

func swap(v []int, a, b int) {
	v[b], v[a] = v[a], v[b]
}
