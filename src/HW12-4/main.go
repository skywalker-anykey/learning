package main

import (
	"fmt"
	"math/rand"
)

func init() {
	//rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println(ar)
	doubleSelectionSort(ar)
	fmt.Println(ar)
}
func doubleSelectionSort(ar []int) {
	minPos := 0
	maxPos := len(ar) - 1

	for minPos <= maxPos {
		minIndex := minPos
		maxIndex := maxPos
		for i := minPos; i <= maxPos; i++ {
			if ar[i] < ar[minIndex] {
				minIndex = i
			}
			if ar[i] > ar[maxIndex] {
				maxIndex = i
			}
		}

		ar[minPos], ar[minIndex] = ar[minIndex], ar[minPos]
		if maxIndex == minPos {
			maxIndex = minIndex
		}
		ar[maxPos], ar[maxIndex] = ar[maxIndex], ar[maxPos]

		minPos++
		maxPos--
	}
}

func selectionSort(ar []int) {
	j := 0

	for range ar {
		find := false
		minIndex := j
		for i := j + 1; i < len(ar); i++ {
			if ar[minIndex] > ar[i] {
				minIndex = i
				find = true
			}
		}
		if find {
			ar[j], ar[minIndex] = ar[minIndex], ar[j]
		}
		j++
	}
}

func selectionSortByMax(ar []int) {
	j := len(ar) - 1

	for j > -1 {
		find := false
		maxIndex := j

		for i := j; i > -1; i-- {
			if ar[i] > ar[maxIndex] {
				maxIndex = i
				find = true
			}
		}
		if find {
			ar[j], ar[maxIndex] = ar[maxIndex], ar[j]
		}

		j--
	}

}
