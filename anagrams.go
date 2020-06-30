package main

import (
	"fmt"
	"encoding/json"
)


func FindAnagramsInList(words []string) [][]string {
	anagramsMap := make(map[string][]string)
	for _, word := range words {
		lettersKey := wordToLettersKey(word)
		if anagrams, ok := anagramsMap[lettersKey]; ok {
			anagramsMap[lettersKey] = append(anagrams, word)
		}else{
			anagramsMap[lettersKey] = []string{word}
		}
	}
	return anagramsMapValues(anagramsMap)
}

func anagramsMapValues(anagramsMap map[string][]string) [][]string {
	anagramsList := make([][]string, 0, len(anagramsMap))
	for _, anagrams := range anagramsMap {
		anagramsList = append(anagramsList, anagrams)
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

func main() {
	fmt.Println("Hello, world.")
}
