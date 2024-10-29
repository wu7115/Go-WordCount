package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type texManager struct {
	filePath string
	data     string
	words    []string
}

func (tm *texManager) dispatch(message string, textFilePath string) {
	if message == "init" {
		tm.readTextFile(textFilePath)
	} else if message == "run" {
		tm.normalize()
	}
}

func (tm *texManager) readTextFile(textFilePath string) {
	textFile, err := os.Open(textFilePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	defer textFile.Close()
	var text strings.Builder
	scanner := bufio.NewScanner(textFile)
	for scanner.Scan() {
		text.WriteString(scanner.Text() + "\n")
	}
	tm.data = text.String()
}

func (tm *texManager) normalize() {
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	tm.words = regex.FindAllString(tm.data, -1)
	for i := range tm.words {
		tm.words[i] = strings.ToLower(tm.words[i])
	}
}

type stopWordManager struct {
	stopWords     map[string]struct{}
	filteredWords []string
}

func (swm *stopWordManager) dispatch(message string, words []string) {
	if message == "init" {
		swm.loadStopWords()
	} else if message == "run" {
		swm.removeStopWords(words)
	}
}

func (swm *stopWordManager) loadStopWords() {
	swm.stopWords = make(map[string]struct{})
	stopFile, err := os.Open("../stop_words.txt")
	if err != nil {
		fmt.Println("Error reading stop file: ", err)
	}
	defer stopFile.Close()
	scanner := bufio.NewScanner(stopFile)
	for scanner.Scan() {
		line := scanner.Text()
		stopwords := strings.Split(line, ",")
		for _, stopword := range stopwords {
			swm.stopWords[stopword] = struct{}{}
		}
	}
}

func (swm *stopWordManager) removeStopWords(words []string) {
	for _, word := range words {
		if _, isStopWord := swm.stopWords[word]; !isStopWord {
			if len(word) > 1 {
				swm.filteredWords = append(swm.filteredWords, word)
			}
		}
	}
}

type frequencyManager struct {
	wordCount map[string]int
}

func (fm *frequencyManager) dispatch(message string, filteredWords []string) {
	if message == "increment" {
		fm.frequencies(filteredWords)
	} else if message == "sort" {
		fm.sortAndPrint()
	}
}

func (fm *frequencyManager) frequencies(filteredWords []string) {
	fm.wordCount = make(map[string]int)
	for _, word := range filteredWords {
		fm.wordCount[word]++
	}
}

func (fm *frequencyManager) sortAndPrint() {
	type wordFreq struct {
		word  string
		count int
	}
	var wordFreqs []wordFreq
	for word, count := range fm.wordCount {
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

type wordFrequencyController struct {
	tm  texManager
	swm stopWordManager
	fm  frequencyManager
}

func (wfc *wordFrequencyController) dispatch(message string) {
	if message == os.Args[1] {
		wfc.init(message)
	} else if message == "run" {
		wfc.run()
	}
}

func (wfc *wordFrequencyController) init(textFilePath string) {
	wfc.tm.dispatch("init", textFilePath)
	wfc.swm.dispatch("init", wfc.tm.words)
}

func (wfc *wordFrequencyController) run() {
	wfc.tm.dispatch("run", "")
	wfc.swm.dispatch("run", wfc.tm.words)
	wfc.fm.dispatch("increment", wfc.swm.filteredWords)
	wfc.fm.dispatch("sort", nil)
}

func main() {
	textFilePath := os.Args[1]
	controller := wordFrequencyController{}
	controller.dispatch(textFilePath)
	controller.dispatch("run")
}
