package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

var data string
var line string
var lowerText string
var words []string
var wordCount = make(map[string]int)
var stop_words []string
var stopWords = make(map[string]struct{})

type wordFreq struct {
	word  string
	count int
}

var wordFreqs []wordFreq

func readFile(pathToFile string) {
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error reading text file: ", err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file contents:", err)
	}
	data = string(text)
}

func toLowerAndCountWords() {
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := regex.FindAllString(line, -1)
		for _, word := range words {
			if len(word) >= 2 {
				if _, isStopWord := stopWords[word]; !isStopWord {
					wordCount[word]++
				}
			}
		}
	}
}

func scanStopWords(pathToFile string) {
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error reading text file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stop_line := scanner.Text()
		stop_words = strings.Split(stop_line, ",")
	}
}

func mapStopWords() {
	for _, n := range stop_words {
		stopWords[n] = struct{}{}
	}
}

func frequencies() {
	for _, n := range words {
		if _, isStopWord := stopWords[n]; !isStopWord {
			wordCount[n]++
		}
	}
}

func sorting() {
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].count > wordFreqs[j].count
	})
}

func main() {
	readFile(os.Args[1])
	scanStopWords(os.Args[2])
	mapStopWords()
	toLowerAndCountWords()
	sorting()

	for i, wf := range wordFreqs {
		if i > 24 {
			break
		}
		fmt.Println(wf.word, " - ", wf.count)
	}
}
