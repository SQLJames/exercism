package scrabble

import (
	"strings"
)

var (
	ones   = []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	twos   = []string{"D", "G"}
	threes = []string{"B", "C", "M", "P"}
	fours  = []string{"F", "H", "V", "W", "Y"}
	fives  = []string{"K"}
	eights = []string{"J", "X"}
	tens   = []string{"Q", "Z"}
)

type (
	letter map[string]int
)

type alphabet struct {
	letters letter
}

func newAlphabet() (alphabet alphabet) {
	alphabet.letters = make(map[string]int, 26)
	for _, v := range ones {
		alphabet.letters[v] = 1
	}
	for _, v := range twos {
		alphabet.letters[v] = 2
	}
	for _, v := range threes {
		alphabet.letters[v] = 3
	}
	for _, v := range fours {
		alphabet.letters[v] = 4
	}
	for _, v := range fives {
		alphabet.letters[v] = 5
	}
	for _, v := range eights {
		alphabet.letters[v] = 8
	}
	for _, v := range tens {
		alphabet.letters[v] = 10
	}
	return alphabet
}

func (a *alphabet) lookupLetterScore(letter string) int {
	return a.letters[letter]
}

func Score(word string) (wordScore int) {
	if word == "" {
		return 0
	}
	wordSlice := strings.Split(strings.ToUpper(word), "")
	a := newAlphabet()
	for _, wordLetter := range wordSlice {
		wordScore += a.lookupLetterScore(wordLetter)
	}
	return wordScore
}
