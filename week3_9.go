package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type fn1 func(map[string]int)
type fn2 func([]string, fn1)
type fn3 func([]string, fn2)
type fn4 func([]string, fn3)
type fn5 func(string, fn4)

func readTextFile(filepath string, f fn5) {
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
	f(text.String(), toLower)
}

func filterChars(text string, f fn4) {
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	words := regex.FindAllString(text, -1)
	f(words, removeStopWords)
}

func toLower(words []string, f fn3) {
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	f(words, frequencies)
}

func removeStopWords(words []string, f fn2) {
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
	f(filteredWords, sortAndPrint)
}

func frequencies(filteredWords []string, f fn1) {
	wordCount := make(map[string]int)
	for _, word := range filteredWords {
		wordCount[word]++
	}
	f(wordCount)
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
	readTextFile(os.Args[1], filterChars)
}
