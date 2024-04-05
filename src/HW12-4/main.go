package main

import (
	"fmt"
)

func init() {
	//rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 10)
	//for i := range ar {
	//	ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	//}
	ar = []int{13, 6, 11, -5, 9, 0}

	fmt.Println(ar)
	doubleSelectionSort(ar)
	fmt.Println(ar)
}
func doubleSelectionSort(ar []int) {
	first := 0
	last := len(ar) - 1
	j := 0

	for first+j < last-j {
		minIndex := first + j
		minFlag, maxFlag := false, false
		maxIndex := last - j

		for i := first + j; i < last-j+1; i++ {
			if ar[i] < ar[minIndex] {
				minIndex = i
				minFlag = true
			}

			if ar[i] > ar[maxIndex] {
				maxIndex = i
				maxFlag = true
			}
		}
		switch {
		//TODO неразобрался с двойным совпадением
		case maxFlag && minFlag:
			ar[first+j], ar[minIndex] = ar[minIndex], ar[first+j]
			ar[minIndex], ar[last-j] = ar[last-j], ar[minIndex]
		case minFlag:
			ar[first+j], ar[minIndex] = ar[minIndex], ar[first+j]
		case maxFlag:
			ar[last-j], ar[maxIndex] = ar[maxIndex], ar[last-j]
		}
		fmt.Println(ar)
		j++
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
