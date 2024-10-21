package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
Суть задания — написать программу, которая считывает из файла список математических выражений, считает результат и записывает в другой файл.

Пример входного файла:
5+4=?
9+3=?
Сегодня прекрасная погода
13+7=?
4-2=?

Пример файла с выводом:
5+4=9
9+3=12
13+7=20
4-2=2

Пожелания к программе:
 -[X] Использовать методы и структуры пакетов ioutils и regexp.
 -[X] Программа должна принимать на вход 2 аргумента: имя входного файла и имя файла для вывода результатов.
 -[X] Если не найден вывод, создать.
 -[X] Если файл вывода существует, очистить перед записью новых результатов.
 -[X] Использовать буферизированную запись результатов.
*/

func main() {
	// -[X] Программа должна принимать на вход 2 аргумента: имя входного файла и имя файла для вывода результатов.
	// Обработка аргументов, если аргументы не указаны используються по заданые умолчанию.
	// --help для получения опций
	inFile := flag.String("in", "./in.txt", "Путь к входящему файлу")
	outFile := flag.String("out", "./out.txt", "Путь к выходящему файлу")

	// Обработка аргументов
	flag.Parse()

	// -[X] Если файл вывода существует, очистить перед записью новых результатов.
	_ = os.Remove(*outFile)
	// -[X] Если не найден вывод, создать.
	fOut, err := os.OpenFile(*outFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer func(fOut *os.File) {
		_ = fOut.Close()
	}(fOut)

	// -[X] Использовать буферизированную запись результатов.
	writer := bufio.NewWriter(fOut)
	if err != nil {
		log.Fatal(err)
	}

	fIn, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`([0-9]+)([+\-*/])([0-9]+)=?`)
	find := re.FindAllStringSubmatch(string(fIn), -1)

	for _, s := range find {
		// Подготовка записи файла в буфер
		_, err = writer.Write([]byte(calc(s[1], s[2], s[3])))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Запись в файл
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func calc(arg1, operation, arg2 string) string {
	a, err := strconv.Atoi(arg1)
	if err != nil {
		log.Fatal(err)
	}

	b, err := strconv.Atoi(arg2)
	if err != nil {
		log.Fatal(err)
	}

	switch operation {
	case "+":
		return fmt.Sprintf("%s+%s=%d\n", arg1, arg2, a+b)
	case "-":
		return fmt.Sprintf("%s-%s=%d\n", arg1, arg2, a-b)
	case "*":
		return fmt.Sprintf("%s*%s=%d\n", arg1, arg2, a*b)
	case "/":
		if b == 0 {
			return ""
		}
		return fmt.Sprintf("%s/%s=%d\n", arg1, arg2, a/b)
	default:
		return ""
	}
}
