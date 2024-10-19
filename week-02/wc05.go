package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	text := strings.ToLower(string(func() []byte { file, _ := os.ReadFile(os.Args[1]); return file }()))
	words := regexp.MustCompile(`[a-z0-9]+{2,}`).FindAllString(text, -1)
	stopWords := strings.Split(string(func() []byte { stopfile, _ := os.ReadFile("../stop_words.txt"); return stopfile }()), ",")
	wordCount := map[string]int{}
	stopWordMap := make(map[string]struct{}, len(stopWords))
	for _, word := range stopWords {
		stopWordMap[word] = struct{}{}
	}
	for _, word := range words {
		if _, isStopWord := stopWordMap[word]; !isStopWord {
			wordCount[word]++
		}
	}
	type wordFreq struct {word  string, count int}
	var wordFreqs []wordFreq
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}
	sort.Slice(wordFreqs, func(i, j int) bool { return wordFreqs[i].count > wordFreqs[j].count })
	for i := 0; i < 25 && i < len(wordFreqs); i++ {
		fmt.Println(wordFreqs[i].word, "-", wordFreqs[i].count)
	}
}
