package day2

import (
	"os"
	"strconv"
	"strings"
)

func ReadGames(path string) []map[string]int {
	data, _ := os.ReadFile(path)
	gamesRecord := string(data)
	games := strings.Split(gamesRecord, "\n")

	gameInfo := make([]map[string]int, 0)

	for _, game := range games {
		buckets := ReadGameLine(game)
		gameInfo = append(gameInfo, buckets)
	}

	return gameInfo
}

func ReadGameLine(game string) map[string]int {
	buckets := make(map[string]int)

	gameInformation := strings.Split(game, ":")[1]
	diceSets := strings.Split(gameInformation, ";")

	for _, set := range diceSets {
		diceCombinations := strings.Split(set, ", ")

		for _, dice := range diceCombinations {
			trimmedDice := strings.Trim(dice, " ")
			colourWithAmount := strings.Split(trimmedDice, " ")
			numOfDice, _ := strconv.Atoi(colourWithAmount[0])

			if buckets[colourWithAmount[1]] < numOfDice {
				buckets[colourWithAmount[1]] = numOfDice
			}
		}
	}

	return buckets
}
