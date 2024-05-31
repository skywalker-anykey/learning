package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	var i int64
	for i = 1000; i < 1015; i++ {
		fmt.Println(hashint64(i))
	}

	fmt.Println(hashstr("abc"))
	fmt.Println(hashstr("cba"))

}

// Напишите свою собственную хеш-функцию, которая будет возвращать хеш (типа uint64) от числа (типа int64), используя остаток от деления на 1000.
func hashint64(val int64) uint64 {
	return uint64(val % 1000)
}

// Напишите свою собственную хеш-функцию, которая будет возвращать хеш (типа uint64) от строки (типа string), используя остаток от деления на 1000.
func hashstr(val string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(val))
	return h.Sum64()
}
