package main

import (
	"bufio"
	"io"
	"os"
)

const dictionaryPath = "dictionary.txt"

var dictionary []string

func loadDictionary() int {
	f, err := os.Open(dictionaryPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var lines int
	for {
		val, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		dictionary = append(dictionary, val)
		lines++
	}
	return lines
}
