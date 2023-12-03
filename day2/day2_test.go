package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day1Part1FirstGameIsPossible(t *testing.T) {
	games := ReadGames("./fixtures/example.txt")
	isPotentialGame := IsPossibleGame(games[0], map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.True(t, isPotentialGame)
}

func Test_Day1Part1SecondGameIsPossible(t *testing.T) {
	games := ReadGames("./fixtures/example.txt")
	isPotentialGame := IsPossibleGame(games[1], map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.True(t, isPotentialGame)
}

func Test_Day1Part1ThirdGameIsNotPossible(t *testing.T) {
	games := ReadGames("./fixtures/example.txt")
	isPotentialGame := IsPossibleGame(games[2], map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.False(t, isPotentialGame)
}

func Test_Day1Part1ExampleHas8(t *testing.T) {
	games := ReadGames("./fixtures/example.txt")
	sum := CalculateGameSum(games, map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.Equal(t, 8, sum)
}

func Test_Day1Part1(t *testing.T) {
	games := ReadGames("./fixtures/part1.txt")
	sum := CalculateGameSum(games, map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.Equal(t, 2317, sum)
}

func Test_Day1Part1FirstGameHas48Power(t *testing.T) {
	games := ReadGames("./fixtures/example.txt")
	totalPower := GetTotalPowerOfSet(games[0])
	assert.Equal(t, 48, totalPower)
}

func Test_Day1Part2(t *testing.T) {
	games := ReadGames("./fixtures/part1.txt")
	sum := CalculateGameSumPart2(games)
	assert.Equal(t, 2317, sum)
}