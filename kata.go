package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	isbn := "048665088x"
	fmt.Println(ValidISBN10(isbn))
}

func ValidISBN10(isbn string) bool {
	var result bool
	anew := strings.Join(strings.Split(isbn, ""), ",")

	var anewSlise []string
	anewSlise = append(anewSlise, anew)

	var ints []int
	anewSlises := strings.Split(anew, ",")

	for _, s := range anewSlises {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	switch {
	case len(isbn) != 10:
		result = false
	default:
		anew := strings.Join(strings.Split(isbn, ""), ",")

		var anewSlise []string
		anewSlise = append(anewSlise, anew)

		var ints []int
		anewSlises := strings.Split(anew, ",")

		for _, s := range anewSlises {
			num, err := strconv.Atoi(s)
			if err == nil {
				ints = append(ints, num)
			}
		}

		if len(isbn) != len(ints) {
			if isbn[len(isbn)-1:] == "X" || isbn[len(isbn)-1:] == "x" {
				ints = append(ints, 10)
			} else {
				result = false
			}

		}

		if len(ints) != 10 {
			result = false
		} else {
			var counter int
			for i := 1; i != 11; i++ {
				counter += i * ints[i-1]
			}

			if counter%11 == 0 {
				result = true
			} else {
				result = false
			}
		}
	}

	return result
}
