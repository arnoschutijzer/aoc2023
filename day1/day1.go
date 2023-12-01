package day1

import (
	"regexp"
	"strconv"
)

func CalculateCalibrationValue(calibrationDocument []string) int {
	total := 0

	for _, line := range calibrationDocument {
		numberAsString := FindNumberInCalibrationDocumentLine(line);
		actualNumber, err := strconv.Atoi(numberAsString);
		if (err != nil) {
			panic("invalid number!")
		}
		total += actualNumber
	}

	return total
}

func FindNumberInCalibrationDocumentLine(input string) string {
	re := regexp.MustCompile(`(\d)+`)
	numbers := re.FindAllString(input, -1)

	if len(numbers) > 1 {
		return numbers[0] + numbers[len(numbers) -1];
	}

	return numbers[0] + numbers[0]
}