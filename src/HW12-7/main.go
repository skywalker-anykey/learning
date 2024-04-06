package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println("Быстрая сортировка")
	fmt.Println("Перед: ", ar)
	fmt.Println("После: ", quickSort(ar))
}

func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	pivot := ar[0]
	var less, greater []int
	for _, num := range ar[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}
