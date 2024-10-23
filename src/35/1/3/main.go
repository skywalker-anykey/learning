package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// dd.mm.YYYY format
var re1 = regexp.MustCompile(`(\d{2}).(\d{2}).(\d{4})`)

// YYYY/mm/dd format
var re2 = regexp.MustCompile(`(\d{4})/(\d{2})/(\d{2})`)

func main() {
	dates := []string{
		"12.09.1978",
		"1990/06/10",
		"08.03.2021",
		"12.04.1986",
		"25 dec 1988",
		"2001/05/25",
	}

	for _, d := range dates {
		year, month, day, err := parseDate(d)
		if err != nil {
			fmt.Println("ERROR!", err, "-", d)
			continue
		}

		fmt.Printf("%04d.%02d.%02d\n", year, month, day)
	}
}

var UndefinedDateFormat = errors.New("undefined date format")

func parseDate(date string) (year, month, day int64, err error) {
	// try dd.mm.YYYY format
	find1 := re1.FindStringSubmatch(date)
	if len(find1) > 0 {
		day = strToInt(find1[1])
		month = strToInt(find1[2])
		year = strToInt(find1[3])
		return
	}

	// try YYYY/mm/dd format
	find2 := re2.FindStringSubmatch(date)
	if len(find2) > 0 {
		day = strToInt(find2[3])
		month = strToInt(find2[2])
		year = strToInt(find2[1])
		return
	}

	// or error
	err = UndefinedDateFormat
	return
}

func strToInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
