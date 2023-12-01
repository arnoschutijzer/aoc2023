package day1

import (
	"fmt"
	"regexp"
	"strconv"
)

func CalculateCalibrationValue(calibrationDocument []string) int {
	total := 0

	for i, line := range calibrationDocument {
		numberAsString := FindNumberInCalibrationDocumentLine(line);
		actualNumber, err := strconv.Atoi(numberAsString);
		fmt.Printf("%d - %d\n", i, actualNumber)
		if (err != nil) {
			panic("invalid number!")
		}
		total += actualNumber
	}

	return total
}

func FindNumberInCalibrationDocumentLine(input string) string {
	re := regexp.MustCompile(`(\d)`)
	numbers := re.FindAllString(input, -1)
	fmt.Printf("%d", len(numbers));

	if len(numbers) > 1 {
		fmt.Printf("%s - %s", numbers[0], numbers[1]);
		return numbers[0] + numbers[len(numbers) -1];
	}

	return numbers[0] + numbers[0]
}