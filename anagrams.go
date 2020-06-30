package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func FindAnagramsInList(words []string, minNumber int) [][]string {
	anagramsMap := make(map[string][]string)
	for _, word := range words {
		lettersKey := wordToLettersKey(word)
		if anagrams, ok := anagramsMap[lettersKey]; ok {
			anagramsMap[lettersKey] = append(anagrams, word)
		} else {
			anagramsMap[lettersKey] = []string{word}
		}
	}
	return anagramsMapToList(anagramsMap, minNumber)
}

func anagramsMapToList(anagramsMap map[string][]string, minNumber int) [][]string {
	anagramsList := make([][]string, 0, len(anagramsMap))
	for _, anagrams := range anagramsMap {
		if len(anagrams) >= minNumber {
			anagramsList = append(anagramsList, anagrams)
		}
	}
	return anagramsList
}

func wordToLettersKey(word string) string {
	letters := wordToLettersMap(word)
	lettersJson, _ := json.Marshal(letters)
	return string(lettersJson)
}

func wordToLettersMap(word string) map[rune]int {
	lettersMap := make(map[rune]int)
	for _, char := range word {
		if nb, ok := lettersMap[char]; ok {
			lettersMap[char] = nb + 1
		} else {
			lettersMap[char] = 1
		}
	}
	return lettersMap
}

func readFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func main() {
	words := readFileLines("/usr/share/dict/words")
	minNumber := 4
	anagramsList := FindAnagramsInList(words, minNumber)
	for _, anagrams := range anagramsList {
		fmt.Println(strings.Join(anagrams, " "))
	}
}
