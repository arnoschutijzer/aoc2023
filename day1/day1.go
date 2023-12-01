package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateCalibrationValue(calibrationDocument []string) int {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9

	total := 0

	for _, line := range calibrationDocument {
		numberAsString := FindNumberInCalibrationDocumentLine(line);
		actualNumber, err := strconv.Atoi(numberAsString);
		if (err != nil) {
			total += m[numberAsString]
		}
		total += actualNumber
	}

	return total
}

var lookup = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

func FindNumberInCalibrationDocumentLine(input string) string {
	firstIndex := -1
	first := -1
	lastIndex := -1
	last := -1

	characters := strings.Split(input, "")
	for index, character := range characters {
		number, err := strconv.Atoi(character)
		if err == nil {
			if first == -1 {
				firstIndex = index;
				first = number
			}
	
			lastIndex = index;
			last = number	
		}
	}

	for index := range characters {
		for key, value := range lookup {
			if strings.HasPrefix(input[index:], key) {
				if (index < firstIndex) {
					firstIndex = index;
					first = value
				}
				if (index > lastIndex) {
					lastIndex = index
					last = value
				}
			}
		}
	}
	return fmt.Sprintf("%d%d", first, last)
}
