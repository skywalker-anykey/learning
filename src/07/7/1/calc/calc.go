package calc

import (
	"log"
)

type calculator struct {
}

// New функция-конструктор для создания экземпляра структуры calculator
func New() calculator {
	return calculator{}
}

//Добавьте неэкспортируемые (приватные) методы для каждой операции (сложения, вычитания, умножения, деления).
//Каждый приватный метод должен принимать на вход 2 числа типа float64 и возвращать значение типа float64.
//(В методе деления чисел должны быть проверка делителя на равенство 0).

// add метод сложения полученых чисел
func (c *calculator) add(x, y float64) float64 {
	return x + y
}

// sub метод вычитания полученых чисел
func (c *calculator) sub(x, y float64) float64 {
	return x - y
}

// mul метод умножения полученых чисел
func (c *calculator) mul(x, y float64) float64 {
	return x * y
}

// div метод деления полученых чисел
func (c *calculator) div(x, y float64) float64 {
	if y == 0 {
		log.Panic("ошибка: попытка деления на ноль")
	}
	return x / y
}

// Calculate метод для структуры calculator.
// В качестве параметров метод должен принимать 2 числа типа float64 и строчный оператор (+, -, *, /).
// Возвращать метод должен значение типа float64.
// Метод Calculate должен в зависимости от переданного оператора вызывать один из приватных методов.
// (Если в качестве оператора передан +, то должен быть вызван приватный метод для сложения чисел).
func (c *calculator) Calculate(x, y float64, operator string) float64 {
	switch operator {
	case "+":
		return c.add(x, y)
	case "-":
		return c.sub(x, y)
	case "*":
		return c.mul(x, y)
	case "/":
		return c.div(x, y)
	default:
		log.Panic("ошибка: неизвестный оператор")
		return 0
	}
}
