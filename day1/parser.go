package day1

import (
	"os"
	"strings"
)

func ReadCalibrationDocument(path string) []string {
	data, _ := os.ReadFile(path)
	calibrationDocument := string(data)
	return strings.Split(calibrationDocument, "\n")
}
