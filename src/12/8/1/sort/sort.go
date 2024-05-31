package sort

func bubbleSort(ar []int) {
	flag := false
	for !flag {
		flag = true
		for i := 1; i < len(ar); i++ {
			if ar[i-1] > ar[i] {
				ar[i-1], ar[i] = ar[i], ar[i-1]
				flag = false
			}
		}
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

func insertionSort(ar []int) {
	// проходим со второго элемента до конца
	for i := 1; i < len(ar); i++ {

		// текущий элемент
		key := ar[i]
		// индекс начального элемента "слева"
		j := i - 1

		// пока слева есть элементы или элемент слева больше текущего
		for j >= 0 && ar[j] > key {
			// сдвигаем элемент
			ar[j+1] = ar[j]
			// следующий элемент "слева"
			j--
		}
		// устанавливаем текущий элемент на его позицию
		ar[j+1] = key
	}
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
