package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	var people = map[string]Man{
		"Джон":   {"Джон", "Вик", 45, "М", 100500},
		"Джеймс": {"Джеймс", "Бонд", 33, "М", 489},
		"Сара":   {"Сара", "Конор", 50, "Ж", 230},
		"Джеки":  {"Джеки", "Чан", 65, "М", 0},
		"Джек":   {"Джек", "Воробей", 53, "М", 30210},
		"Гарри":  {"Гарри", "Поттер", 18, "М", 1},
		"Алиса":  {"Алиса", "Селезнёва", 14, "Ж", 5},
	}

	if len(people) == 0 {
		fmt.Println("В базе данных нет людей")
		return
	}

	suspects := make([]string, 0, len(people))
	for s := range people {
		suspects = append(suspects, s)
	}

	topOne := Man{}
	findSomeOne := false
	for _, suspect := range suspects {
		if people[suspect].Crimes > topOne.Crimes {
			findSomeOne = true
			topOne = people[suspect]
		}
	}

	if findSomeOne {
		fmt.Println("Подозреваемый: ", topOne.Name, topOne.LastName)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}
