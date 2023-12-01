package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day1Part1ExampleHas142(t *testing.T) {
	calibrationDocument := ReadCalibrationDocument("./fixtures/example.txt")
	calibrationValue := CalculateCalibrationValue(calibrationDocument)
	assert.Equal(t, 142, calibrationValue)
}

func Test_FindsFirstAndLastInteger(t *testing.T) {
	var numbers string

	numbers = FindNumberInCalibrationDocumentLine("1abc2")
	assert.Equal(t, "12", numbers)

	numbers = FindNumberInCalibrationDocumentLine("pqr3stu8vwx")
	assert.Equal(t, "38", numbers)

	numbers = FindNumberInCalibrationDocumentLine("a1b2c3d4e5f")
	assert.Equal(t, "15", numbers)

	numbers = FindNumberInCalibrationDocumentLine("treb7uchet")
	assert.Equal(t, "77", numbers)
}

