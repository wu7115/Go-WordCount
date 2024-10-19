package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

func mergeSort(wordFreqs []wordFreq) []wordFreq {
	n := len(wordFreqs)
	temp := make([]wordFreq, n)
	for size := 1; size < n; size *= 2 {
		for leftStart := 0; leftStart < n; leftStart += 2 * size {
			mid := leftStart + size
			rightStart := mid
			end := leftStart + 2*size
			if mid > n {
				mid = n
			}
			if rightStart > n {
				rightStart = n
			}
			if end > n {
				end = n
			}
			merge(wordFreqs, temp, leftStart, mid, rightStart, end)
		}
		copy(wordFreqs, temp)
	}
	return wordFreqs
}

func merge(wordFreqs, temp []wordFreq, leftStart, mid, rightStart, end int) {
	i, j, k := leftStart, rightStart, leftStart
	for i < mid && j < end {
		if wordFreqs[i].count >= wordFreqs[j].count {
			temp[k] = wordFreqs[i]
			i++
		} else {
			temp[k] = wordFreqs[j]
			j++
		}
		k++
	}
	for i < mid {
		temp[k] = wordFreqs[i]
		i++
		k++
	}
	for j < end {
		temp[k] = wordFreqs[j]
		j++
		k++
	}
}

type wordFreq struct {
	word  string
	count int
}

var wordFreqs []wordFreq

func main() {
	textFilePath := os.Args[1]
	stopWordsPath := "../stop_words.txt"
	file, err := os.Open(textFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	stopWords := StopWordsMap(stopWordsPath)
	wordCount := make(map[string]int)
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	scanner := bufio.NewScanner(file)
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
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}
	wordFreqs = mergeSort(wordFreqs)
	for i, wf := range wordFreqs {
		if i > 24 {
			break
		}
		fmt.Println(wf.word, " - ", wf.count)
	}
}
