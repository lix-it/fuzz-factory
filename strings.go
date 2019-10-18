package main

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func substituteRandomCharacter(input string) string {
	outputBytes := []byte(input)
	// get a random character from the input string
	subCharIndex := rand.Intn(len(outputBytes))
	// generate a random character & insert
	outputBytes[subCharIndex] = letters[rand.Intn(len(letters))]

	return string(outputBytes)
}

func addRandomCharacter(input string) string {
	outputBytes := []byte(input)
	outputBytes = append(outputBytes, letters[rand.Intn(len(letters))])
	return string(outputBytes)
}

func deleteRandomCharacter(input string) string {
	outputBytes := []byte(input)
	char := rand.Intn(len(outputBytes))
	// delete
	outputBytes = append(outputBytes[:char], outputBytes[char+1:]...)
	return string(outputBytes)
}
