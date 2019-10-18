package main

import (
	"math/rand"
	"strings"
)

// pick a random word from the input and replace with a dictionary word
func substituteWordRandom(input string) string {
	words := getWords(input)
	repIndex := rand.Intn(len(words))
	words[repIndex] = dictionary[rand.Intn(len(dictionary))]
	return strings.Join(words, " ")
}

func addWordRandom(input string) string {
	words := getWords(input)
	// insert into random place in the output string
	randIndex := rand.Intn(len(words))
	newWord := dictionary[rand.Intn(len(dictionary))]
	words = append(words[:randIndex], append([]string{newWord}, words[randIndex:]...)...)

	return strings.Join(words, " ")
}

func deleteWordRandom(input string) string {
	words := getWords(input)
	randIndex := rand.Intn(len(words))
	// delete index
	words = append(words[:randIndex], words[randIndex+1:]...)
	return strings.Join(words, " ")
}

func getWords(input string) []string {
	return strings.Split(input, " ")
}
