package main

import (
	"fmt"
	"sync"
)

const (
	// numGoroutines - Количество используемых горутин
	numGoroutines = 10
	// goal - Цель конечного значения счетчика
	goal = 1000
)

type Counter struct {
	value  int
	ch     chan int
	goal   int
	goalCh chan bool
}

func NewCounter() *Counter {
	ch := make(chan int)
	goalCh := make(chan bool)
	value := 0

	return &Counter{
		value: value,
		ch:    ch,

		goalCh: goalCh,
		goal:   goal,
	}
}

func (c *Counter) Start() {
	go func() {
		for step := range c.ch {
			if c.value < c.goal {
				c.value += step
				c.goalCh <- false
			} else {
				c.goalCh <- true
			}
		}
	}()
}

func (c *Counter) Add(step int) bool {
	c.ch <- step
	if <-c.goalCh == true {
		return true
	}
	return false
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Stop() {
	close(c.ch)
	close(c.goalCh)
}

func main() {
	wg := &sync.WaitGroup{}
	c := NewCounter()

	c.Start()

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				isGoal := c.Add(1)
				if isGoal {
					return
				}
			}
		}()
	}

	wg.Wait()

	c.Stop()

	fmt.Println(c.Value())
}

// Напишите код, в котором несколько горутин увеличивают значение целочисленного счётчика и синхронизируют свою работу
// через канал. Нужно предусмотреть возможность настройки количества используемых горутин и конечного значения счётчика,
// до которого его следует увеличивать.
//
// - Попробуйте реализовать счётчик с элементами ООП (в виде структуры и методов структуры).
// - Попробуйте реализовать динамическую проверку достижения счётчиком нужного значения.
