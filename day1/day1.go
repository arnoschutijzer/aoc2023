package day1

import (
	"fmt"
	"regexp"
	"strconv"
)

func CalculateCalibrationValue(calibrationDocument []string) int {
	total := 0

	for _, line := range calibrationDocument {
		numberAsString := FindNumberInCalibrationDocumentLine(line);
		actualNumber, _ := strconv.Atoi(numberAsString);
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
	re := regexp.MustCompile(`(\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
	first := re.FindAllString(input, -1)[0]

	// hehe
	reversedRegex := regexp.MustCompile(`(\d)|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)`)
	reversed := reverse(input)
	last := reverse(reversedRegex.FindAllString(reversed, -1)[0])

	firstNumber, err := strconv.Atoi(first)
	if err != nil {
		firstNumber = lookup[first]
	}
	lastNumber, err := strconv.Atoi(last)
	if err != nil {
		lastNumber = lookup[last]
	}

	return fmt.Sprintf("%d%d", firstNumber, lastNumber)
}

func reverse(s string) string {
    chars := []rune(s)
    for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
        chars[i], chars[j] = chars[j], chars[i]
    }
    return string(chars)
}