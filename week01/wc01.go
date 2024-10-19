package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// load the stop words and return a map of them
func StopWordsMap(filepath string) map[string]struct{} {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	stopWords := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// split the words by comma
		words := strings.Split(line, ",")

		for _, word := range words {
			stopWords[word] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return stopWords
}

// count the word frequency in the given text file
func main() {
	textFilePath := os.Args[1]
	stopWordsPath := os.Args[2]

	file, err := os.Open(textFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	stopWords := StopWordsMap(stopWordsPath)

	// create a map for storing words and their frequency
	wordCount := make(map[string]int)
	// filter out texts that are not alphabets
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)

	// count word frequency line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := regex.FindAllString(line, -1)
		for _, word := range words {
			if len(word) >= 2 {
				// only count words that are not in stop words
				if _, isStopWord := stopWords[word]; !isStopWord {
					wordCount[word]++
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// create a data structure for storing pairs of {word, count}
	type wordFreq struct {
		word  string
		count int
	}
	// append all the apirs to the slice for sorting later
	var wordFreqs []wordFreq
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}

	// sort the slice in non-decreasing order of count
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].count > wordFreqs[j].count
	})

	// only print out the 25 most frequency terms
	for i, wf := range wordFreqs {
		if i > 24 {
			break
		}
		fmt.Println(wf.word, " - ", wf.count)
	}
}
