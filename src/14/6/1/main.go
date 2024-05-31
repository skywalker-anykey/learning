package main

import "fmt"

func main() {
	// Реализуйте программу, которая будет находить общие элементы в двух массивах и использовать map для этой цели.
	// Программа должна работать таким образом:
	// Вы запускаете программу из консоли: 'go run main.go' или './intersection'.

	var arrLen1, arrLen2 int

	// Программа просит ввести размер первого массива (ожидается ввод числа и нажатие каретки (кл. enter)).
	fmt.Println("Enter first array size:")
	fmt.Scanln(&arrLen1)

	// Программа просит ввести размер второго массива (ожидается ввод числа и нажатие каретки (кл. enter)).
	fmt.Println("Enter second array size:")
	fmt.Scanln(&arrLen2)

	// Программа просит ввести элементы первого массива, используя перенос строки как разделитель.
	fmt.Println("Enter first array:")
	arr1 := make(map[string]struct{}, arrLen1)
	for i := 0; i < arrLen1; i++ {
		var k string
		fmt.Scanln(&k)
		arr1[k] = struct{}{}
	}

	// Когда пользователь заполнит первый массив, программа запрашивает аналогичным образом ввод элементов второго массива.
	fmt.Println("Enter second array:")
	arr2 := make(map[string]struct{}, arrLen2)
	for i := 0; i < arrLen2; i++ {
		var k string
		fmt.Scanln(&k)
		arr2[k] = struct{}{}
	}

	// Программа выводит список одинаковых элементов.
	var out []string
	for k := range arr1 {
		if _, ok := arr2[k]; ok {
			out = append(out, k)
		}
	}
	fmt.Println(out)
}
