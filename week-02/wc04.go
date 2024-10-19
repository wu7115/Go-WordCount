package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func readTextFile(filepath string) string {
	textFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	defer textFile.Close()
	var text strings.Builder
	scanner := bufio.NewScanner(textFile)
	for scanner.Scan() {
		text.WriteString(scanner.Text() + "\n")
	}
	return text.String()
}

func filterChars(text string) []string {
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	words := regex.FindAllString(text, -1)
	return words
}

func toLower(words []string) []string {
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func removeStopWords(words []string) []string {
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		fmt.Println("Error reading stop file: ", err)
	}
	defer file.Close()
	stopWords := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stopwords := strings.Split(line, ",")
		for _, stopword := range stopwords {
			stopWords[stopword] = struct{}{}
		}
	}
	filteredWords := []string{}
	for _, word := range words {
		if _, isStopWord := stopWords[word]; !isStopWord {
			if len(word) > 1 {
				filteredWords = append(filteredWords, word)
			}
		}
	}
	return filteredWords
}

func frequencies(filteredWords []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range filteredWords {
		wordCount[word]++
	}
	return wordCount
}

func sortAndPrint(wordCount map[string]int) {
	type wordFreq struct {
		word  string
		count int
	}
	var wordFreqs []wordFreq
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].count > wordFreqs[j].count
	})
	for i, wf := range wordFreqs {
		if i > 24 {
			break
		}
		fmt.Println(wf.word, " - ", wf.count)
	}
}

func main() {
	sortAndPrint(frequencies(removeStopWords(toLower(filterChars(readTextFile(os.Args[1]))))))
}
