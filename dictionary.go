package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"path/filepath"
)

var dictionaryPath = flag.String("d", "dictionary.txt", "dictionary file for where you want to  ")

var dictionary []string

func loadDictionary() int {
	absDictPath, err := filepath.Abs(*dictionaryPath)
	if err != nil {
		panic(err)
	}
	f, err := os.Open(absDictPath)
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
