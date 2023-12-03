package day2

func CalculateGameSum(games []map[string]int, sets map[string]int) int {
	total := 0

	for index, game := range games {
		if IsPossibleGame(game, sets) {
			total += index + 1
		}
	}

	return total
}

func IsPossibleGame(game map[string]int, sets map[string]int) bool {
	for colour, amount := range sets {
		if game[colour] > amount {
			return false
		}
	}
	return true
}
