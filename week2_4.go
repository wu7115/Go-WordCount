package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	stopWordFile, err := os.Open(os.Args[2])
	if err != nil {
		return
	}
	defer stopWordFile.Close()

	var stopWords []string
	stopScanner := bufio.NewScanner(stopWordFile)
	for stopScanner.Scan() {
		line := stopScanner.Text()
		words := strings.Split(line, ",")
		for _, word := range words {
			stopWords = append(stopWords, word)
		}
	}

	textFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer textFile.Close()
	text, err := io.ReadAll(textFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lowerText := strings.ToLower(string(text))
	startChar := -1
	type wordFreq struct {
		word  string
		count int
	}
	var wordFreqs []wordFreq
	found := false
	for i, char := range lowerText {
		isStopWord := false
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if startChar == -1 {
				startChar = i
			}
		} else {
			if startChar != -1 {
				word := string(lowerText[startChar:i])
				startChar = -1
				for n := range stopWords {
					if word == stopWords[n] {
						isStopWord = true
					}
				}
				if !isStopWord {
					if len(wordFreqs) > 1 {
						for n := range wordFreqs {
							if word == wordFreqs[n].word {
								found = true
								wordFreqs[n].count++
								for k := n; k >= 1; k-- {
									if wordFreqs[k].count > wordFreqs[k-1].count {
										wordFreqs[k-1], wordFreqs[k] = wordFreqs[k], wordFreqs[k-1]
									}
								}
							}
						}
						if !found && len(word) > 1 {
							wordFreqs = append(wordFreqs, wordFreq{word, 1})
						}
						found = false
					} else if len(word) > 1 {
						wordFreqs = append(wordFreqs, wordFreq{word, 1})
					}
				}
			}
		}
	}
	for n := 0; n < 25; n++ {
		fmt.Println(wordFreqs[n].word, " - ", wordFreqs[n].count)
	}
}

// func main() {
// 	textFilePath := os.Args[1]
// 	stopWordsPath := os.Args[2]

// 	textFile, err := os.Open(textFilePath)
// 	if err != nil {
// 		return
// 	}
// 	defer textFile.Close()

// 	stopWordFile, err := os.Open(stopWordsPath)
// 	if err != nil {
// 		return
// 	}
// 	defer stopWordFile.Close()

// 	var stopWords []string
// 	stopScanner := bufio.NewScanner(stopWordFile)
// 	for stopScanner.Scan() {
// 		line := stopScanner.Text()
// 		words := strings.Split(line, ",")
// 		for _, word := range words {
// 			stopWords = append(stopWords, word)
// 		}
// 	}

// 	type wordFreq struct {
// 		word  string
// 		count int
// 	}
// 	var wordFreqs []wordFreq

// 	textScanner := bufio.NewScanner(textFile)
// 	for textScanner.Scan() {
// 		i := 0
// 		start_char := -1
// 		line := strings.ToLower(textScanner.Text())
// 		for _, char := range line {
// 			if start_char == -1 {
// 				if unicode.IsLetter(char) || unicode.IsNumber(char) {
// 					start_char = i
// 				}
// 			} else {
// 				if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
// 					found := false
// 					word := line[start_char:i]
// 					isStopWord := false
// 					for _, stopWord := range stopWords {
// 						if word == stopWord {
// 							isStopWord = true
// 							break
// 						}
// 					}
// 					if !isStopWord {
// 						pair_index := 0
// 						for n := range wordFreqs {
// 							if wordFreqs[n].word == word {
// 								wordFreqs[n].count++
// 								found = true
// 								break
// 							}
// 							pair_index++
// 						}
// 						if !found {
// 							wordFreqs = append(wordFreqs, wordFreq{word, 1})
// 						} else if len(wordFreqs) > 1 {
// 							for n := pair_index - 1; n >= 0; n-- {
// 								if wordFreqs[pair_index].count > wordFreqs[n].count {
// 									wordFreqs[n], wordFreqs[pair_index] = wordFreqs[pair_index], wordFreqs[n]
// 									pair_index = n
// 								}
// 							}
// 						}
// 					}
// 					start_char = -1
// 				}
// 			}
// 			i++
// 		}
// 	}
// 	for tf := 0; tf < 24; tf++ {
// 		fmt.Println(wordFreqs[tf].word, " - ", wordFreqs[tf].count)
// 	}
// }
