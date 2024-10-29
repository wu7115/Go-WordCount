package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var dataStorageObj = map[string]interface{}{
	"data": []string{},
}

var stopWordsObj = map[string]interface{}{
	"stop_words": make(map[string]struct{}),
}

var wordFreqsObj = map[string]interface{}{
	"freqs": map[string]int{},
}

var printObj = map[string]interface{}{} // Define output object

func normalize(textFilePath string) {
	textFile, err := os.Open(textFilePath)
	if err != nil {
		fmt.Println("Error: File", textFilePath, "not found.")
	}
	defer textFile.Close()

	scanner := bufio.NewScanner(textFile)
	var dataStr strings.Builder
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)

	for scanner.Scan() {
		line := scanner.Text()
		words := regex.FindAllString(line, -1)
		for _, word := range words {
			dataStr.WriteString(strings.ToLower(word) + " ")
		}
	}

	dataStorageObj["data"] = strings.Fields(dataStr.String())
}

func loadStopWords() {
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		fmt.Println("Error: Stop words file not found.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stopwords := strings.Split(line, ",")
		for _, stopword := range stopwords {
			stopWordsObj["stop_words"].(map[string]struct{})[stopword] = struct{}{}
		}
	}
}

func wordCount(w string) {
	freqs := wordFreqsObj["freqs"].(map[string]int)
	freqs[w]++
}

func sortFrequencies() [][2]interface{} {
	freqs := wordFreqsObj["freqs"].(map[string]int)
	sortedFreqs := make([][2]interface{}, 0, len(freqs))

	for w, c := range freqs {
		sortedFreqs = append(sortedFreqs, [2]interface{}{w, c})
	}

	sort.Slice(sortedFreqs, func(i, j int) bool {
		return sortedFreqs[i][1].(int) > sortedFreqs[j][1].(int)
	})

	return sortedFreqs
}

func printOutput() {
	wordFreqs := wordFreqsObj["sort"].(func() [][2]interface{})()
	for i := 0; i < 25 && i < len(wordFreqs); i++ {
		fmt.Printf("%s - %d\n", wordFreqs[i][0].(string), wordFreqs[i][1].(int))
	}
}

func main() {
	dataStorageObj["init"] = func(textFilePath string) { normalize(textFilePath) }
	dataStorageObj["words"] = func() []string { return dataStorageObj["data"].([]string) }

	stopWordsObj["init"] = func() { loadStopWords() }
	stopWordsObj["is_stop_word"] = func(word string) bool {
		_, isStopWord := stopWordsObj["stop_words"].(map[string]struct{})[word]
		return isStopWord
	}

	wordFreqsObj["word_count"] = func(w string) { wordCount(w) }
	wordFreqsObj["sort"] = func() [][2]interface{} {
		return sortFrequencies()
	}

	printObj["print"] = func() { printOutput() } // Add output function

	dataStorageObj["init"].(func(string))(os.Args[1])
	stopWordsObj["init"].(func())()

	filteredWords := []string{}
	words := dataStorageObj["words"].(func() []string)()
	for _, word := range words {
		if !stopWordsObj["is_stop_word"].(func(string) bool)(word) {
			if len(word) > 1 {
				filteredWords = append(filteredWords, word)
				wordFreqsObj["word_count"].(func(string))(word)
			}
		}
	}

	printObj["print"].(func())() // Call output function
}
