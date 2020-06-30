package main

import "testing"
import "reflect"

func TestFindAnagramsInList(t *testing.T) {
	words := []string{"ab", "ba", "z", "zz"}

	anagrams := FindAnagramsInList(words, 1)

	expected := [][]string{[]string{"ab", "ba"}, []string{"z"}, []string{"zz"}}
	if !reflect.DeepEqual(anagrams, expected) {
		t.Errorf("Expected %s, got%s", expected, anagrams)
	}
}

func TestFindAnagramsInListWithMinNumber(t *testing.T) {
	words := []string{"ab", "ba", "z"}

	anagrams := FindAnagramsInList(words, 2)

	expected := [][]string{[]string{"ab", "ba"}}
	if !reflect.DeepEqual(anagrams, expected) {
		t.Errorf("Expected %s, got%s", expected, anagrams)
	}
}
