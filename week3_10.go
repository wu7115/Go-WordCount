package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Result[T any] struct {
	Value T
	Err   error
}

func Bind[A, B any](r Result[A], f func(A) Result[B]) Result[B] {
	if r.Err != nil {
		return Result[B]{Err: r.Err}
	}
	return f(r.Value)
}

func Success[T any](val T) Result[T] {
	return Result[T]{Value: val}
}

func readTextFile(filepath string) Result[string] {
	textFile, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer textFile.Close()

	var text strings.Builder
	scanner := bufio.NewScanner(textFile)
	for scanner.Scan() {
		text.WriteString(scanner.Text() + "\n")
	}
	return Success(text.String())
}

func filterChars(text string) Result[[]string] {
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	words := regex.FindAllString(text, -1)
	return Success(words)
}

func toLower(words []string) Result[[]string] {
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return Success(words)
}

func removeStopWords(words []string) Result[[]string] {
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		os.Exit(1)
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
	return Success(filteredWords)
}

func frequencies(filteredWords []string) Result[map[string]int] {
	wordCount := make(map[string]int)
	for _, word := range filteredWords {
		wordCount[word]++
	}
	return Success(wordCount)
}

func sortAndPrint(wordCount map[string]int) Result[bool] {
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
	return Success(true)
}

func main() {
	result := Bind(
		Bind(
			Bind(
				Bind(
					Bind(readTextFile(os.Args[1]), filterChars),
					toLower),
				removeStopWords),
			frequencies),
		sortAndPrint)

	if result.Err != nil {
		fmt.Println("Error:", result.Err)
	}
}
