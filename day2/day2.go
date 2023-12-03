package day2

func CalculateGameSum(games []map[string]int, sets map[string]int) int {
	total := 0

	for index, game := range games {
		if isPossibleGame(game, sets) {
			total += index + 1
		}
	}

	return total
}

func CalculateGameSumPart2(games []map[string]int) int {
	total := 0

	for _, game := range games {
		total += GetTotalPowerOfSet(game)
	}

	return total
}

func GetTotalPowerOfSet(game map[string]int) int {
	total := 1

	for _, amount := range game {
		total = total * amount
	}

	return total
}

func isPossibleGame(game map[string]int, sets map[string]int) bool {
	for colour, amount := range sets {
		if game[colour] > amount {
			return false
		}
	}
	return true
}
