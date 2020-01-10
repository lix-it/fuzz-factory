package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Record struct {
	record []string
	index  int
}

var hasHeaders = flag.Bool("headers", false, "input file has first row as headers")

func main() {
	timeNow := time.Now()
	defer func() {
		fmt.Printf("Program took %v\n", time.Since(timeNow))
	}()
	flag.Parse()
	outputFilePath := "output.csv"
	filePath := flag.Arg(0)
	if flag.Arg(1) != "" {
		outputFilePath = flag.Arg(1)
	}
	fuzzFile(filePath, outputFilePath, *hasHeaders)
}

func fuzzFile(filePath string, outputFilePath string, hasHeaders bool) {
	// set up random seeding
	rand.Seed(time.Now().UnixNano())
	// load the input file
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}
	f, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	outputAbsPath, err := filepath.Abs(outputFilePath)
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create(outputAbsPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	// write to output
	csvWriter := csv.NewWriter(outputFile)
	defer csvWriter.Flush()

	// set up dictionary
	loadDictionary()
	fmt.Printf("Dictionary has %v words\n", len(dictionary))
	csvReader := csv.NewReader(f)
	// split CPU-bound fuzzing jobs across goroutines so we can load the next
	// csv record
	done := make(chan bool)
	newRecordChan := make(chan Record)
	var wg sync.WaitGroup
	// collector function
	go func() {
		waitlist := make(map[int]Record)
		var lastWritten int
		for {
			select {
			case <-done:
				return
			case newRecord := <-newRecordChan:
				// add results to output dataset
				// make sure dataset is in order using a waitlist
				waitlist[newRecord.index] = newRecord

				// pop the next ones off the waitlist if they are present
				for {
					val, ok := waitlist[lastWritten]
					if !ok {
						break
					}
					err := csvWriter.Write(val.record)
					if err != nil {
						panic(err)
					}
					lastWritten++
				}
			}
		}
	}()
	var currentRowNum int
	for {
		// do not put this in a goroutine as reads need to be sequential
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		r := Record{record: record, index: currentRowNum}
		currentRowNum++
		// insert header row in without fuzzing
		if hasHeaders && currentRowNum == 1 {
			newRecordChan <- r
			continue
		}
		wg.Add(1)
		go fuzzRecord(r, newRecordChan, &wg)
	}
	wg.Wait()
	done <- true
}

func fuzzRecord(record Record, resultChan chan Record, wg *sync.WaitGroup) {
	defer wg.Done()
	newRecord := record.record
	// determine which function classes we should run
	shouldTypo := rand.Intn(2)

	shouldAbbreviate := rand.Intn(2)
	// determine which columns we should fuzz
	shouldFuzzColumns := []int{}
	for range newRecord {
		shouldFuzzColumns = append(shouldFuzzColumns, rand.Intn(2))
	}
	// TODO: splitting it up like this is actually quite inefficient...
	if shouldTypo == 1 {
		newRecord = generateTypo(newRecord, shouldFuzzColumns)
	}
	if shouldAbbreviate == 1 {
		newRecord = generateAbbreviations(newRecord, shouldFuzzColumns)
	}
	resultChan <- Record{
		record: newRecord,
		index:  record.index,
	}
}

func generateTypo(record []string, columnsToFuzz []int) []string {
	newRecord := record
	for index, shouldFuzzCell := range columnsToFuzz {
		if shouldFuzzCell == 1 {
			// don't fuzz if null
			if newRecord[index] == "" {
				continue
			}
			// shouldCharacterSubstitute
			if rand.Intn(2) == 1 {
				newRecord[index] = substituteRandomCharacter(newRecord[index])
			}

			// shouldCharacterAddition
			if rand.Intn(2) == 1 {
				newRecord[index] = addRandomCharacter(newRecord[index])
			}

			// shouldCharacterDelete
			if rand.Intn(2) == 1 {
				newRecord[index] = deleteRandomCharacter(newRecord[index])
			}

			// shouldWordSubstitute
			if rand.Intn(2) == 1 {
				newRecord[index] = substituteWordRandom(newRecord[index])
			}
			// shouldWordAddtion
			if rand.Intn(2) == 1 {
				newRecord[index] = addWordRandom(newRecord[index])
			}
			// shouldWordDeletion
			if rand.Intn(2) == 1 {
				newRecord[index] = deleteWordRandom(newRecord[index])
			}
		}
	}
	return newRecord
}

func generateAbbreviations(record []string, columnsToAbbr []int) []string {
	newRecord := record
	for index, shouldAbbrCell := range columnsToAbbr {
		if shouldAbbrCell == 1 {
			if newRecord[index] == "" {
				// don't abbreviate null values
				continue
			}
			// shouldAbbreviate
			finalLength := len(newRecord[index]) - rand.Intn(len(newRecord[index]))
			if rand.Intn(2) == 1 {
				newRecord[index] = abbreviateString(newRecord[index], finalLength)
			}
		}
	}

	return newRecord
}
