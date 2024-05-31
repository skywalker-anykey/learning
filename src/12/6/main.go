package main

import (
	"fmt"
	"math/rand"
)

func init() {
	//rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 51)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println("Перед: ", ar)
	ar = mergeSort(ar)
	fmt.Println("После: ", ar)
}

func mergeSort(ar []int) []int {
	// выходим из рекурсии если делить нечего
	if len(ar) < 2 {
		return ar
	}
	// находим индекс середины
	mid := len(ar) / 2
	// делим на на части и отправляем их в рекурсию
	left := mergeSort(ar[:mid])
	right := mergeSort(ar[mid:])

	// возвращаем результат слияния (сортировки)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	// пока слева или справа есть элементы
	for len(left) > 0 && len(right) > 0 {
		// если левый первый элемент меньше(равен) правого первого
		if left[0] <= right[0] {
			// в результат добавляем его первым
			merged = append(merged, left[0])
			// убираем первый элемент из левой части
			left = left[1:]
		} else {
			// в результат добавляем правый первый
			merged = append(merged, right[0])
			// убираем первый элемент из правой части
			right = right[1:]
		}
	}

	// добавляем оставшиеся части
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}
