package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 50
	array := make([]int, n)

	for i := 0; i < n; i++ {
		array[i] = rand.Intn(100) + 1 // Генерирует число от 1 до 100
	}

	fmt.Println("Случайный массив:", array)

	array = TimSort(array)
	fmt.Println("отсортированный массив:", array)

}
