package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordFrequencyFramework struct {
	loadHandlers   []func(string)
	doworkHandlers []func()
	endHandlers    []func()
}

func (wf *WordFrequencyFramework) RegisterForLoadEvent(handler func(string)) {
	wf.loadHandlers = append(wf.loadHandlers, handler)
}

func (wf *WordFrequencyFramework) RegisterForDoWorkEvent(handler func()) {
	wf.doworkHandlers = append(wf.doworkHandlers, handler)
}

func (wf *WordFrequencyFramework) RegisterForEndEvent(handler func()) {
	wf.endHandlers = append(wf.endHandlers, handler)
}

func (wf *WordFrequencyFramework) Run(filePath string) {
	for _, h := range wf.loadHandlers {
		h(filePath)
	}
	for _, h := range wf.doworkHandlers {
		h()
	}
	for _, h := range wf.endHandlers {
		h()
	}
}

type DataStorage struct {
	data           string
	stopWordFilter *StopWordFilter
	wordHandlers   []func(string)
}

func NewDataStorage(wf *WordFrequencyFramework, swf *StopWordFilter) *DataStorage {
	ds := &DataStorage{stopWordFilter: swf}
	wf.RegisterForLoadEvent(ds.load)
	wf.RegisterForDoWorkEvent(ds.produceWords)
	return ds
}

func (ds *DataStorage) load(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: File not found.")
		return
	}
	defer file.Close()

	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	var sb strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := regex.FindAllString(line, -1)
		for _, word := range words {
			if len(word) > 1 {
				sb.WriteString(strings.ToLower(word) + " ")
			}
		}
	}

	ds.data = sb.String()
}

func (ds *DataStorage) produceWords() {
	for _, word := range strings.Fields(ds.data) {
		if !ds.stopWordFilter.IsStopWord(word) {
			for _, h := range ds.wordHandlers {
				h(word)
			}
		}
	}
}

func (ds *DataStorage) RegisterForWordEvent(handler func(string)) {
	ds.wordHandlers = append(ds.wordHandlers, handler)
}

type StopWordFilter struct {
	stopWords map[string]struct{}
}

func NewStopWordFilter(wf *WordFrequencyFramework) *StopWordFilter {
	swf := &StopWordFilter{stopWords: make(map[string]struct{})}
	wf.RegisterForLoadEvent(swf.load)
	return swf
}

func (swf *StopWordFilter) load(filePath string) {
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		fmt.Println("Error: Stop words file not found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ",")
		for _, word := range words {
			swf.stopWords[strings.TrimSpace(word)] = struct{}{}
		}
	}
}

func (swf *StopWordFilter) IsStopWord(word string) bool {
	_, exists := swf.stopWords[word]
	return exists
}

type WordFrequencyCounter struct {
	wordFreqs map[string]int
}

func NewWordFrequencyCounter(wf *WordFrequencyFramework, ds *DataStorage) *WordFrequencyCounter {
	wfc := &WordFrequencyCounter{wordFreqs: make(map[string]int)}
	ds.RegisterForWordEvent(wfc.incrementCount)
	wf.RegisterForEndEvent(wfc.printFreqs)
	wf.RegisterForEndEvent(wfc.printWordsWithZ)
	return wfc
}

func (wfc *WordFrequencyCounter) incrementCount(word string) {
	wfc.wordFreqs[word]++
}

func (wfc *WordFrequencyCounter) printFreqs() {
	type kv struct {
		Key   string
		Value int
	}
	var sortedFreqs []kv
	for k, v := range wfc.wordFreqs {
		sortedFreqs = append(sortedFreqs, kv{k, v})
	}
	sort.Slice(sortedFreqs, func(i, j int) bool {
		return sortedFreqs[i].Value > sortedFreqs[j].Value
	})
	for i, kv := range sortedFreqs {
		if i >= 25 {
			break
		}
		fmt.Printf("%s - %d\n", kv.Key, kv.Value)
	}
}

func (wfc *WordFrequencyCounter) printWordsWithZ() {
	fmt.Println("\nWords in 'z':")
	for word := range wfc.wordFreqs {
		if strings.Contains(word, "z") {
			fmt.Println(word)
		}
	}
}

func main() {
	wf := &WordFrequencyFramework{}
	stopWordFilter := NewStopWordFilter(wf)
	dataStorage := NewDataStorage(wf, stopWordFilter)
	NewWordFrequencyCounter(wf, dataStorage)

	wf.Run(os.Args[1])
}