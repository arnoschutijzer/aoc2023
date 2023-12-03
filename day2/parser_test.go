package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parses1Game(t *testing.T) {
	assert.Equal(t, 6, readGameLine("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")["blue"])
	assert.Equal(t, 3, readGameLine("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")["green"])
	assert.Equal(t, 20, readGameLine("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")["red"])
}

func Test_ParsesEntireRecord(t *testing.T) {
	games := ReadGames("./fixtures/example.txt")
	assert.Equal(t, map[string]int{
		"blue":  6,
		"green": 2,
		"red":   4,
	}, games[0])
}
