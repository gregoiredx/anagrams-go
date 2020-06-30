package main

import "testing"
import "reflect"

func TestFindAnagramsInList(t *testing.T) {
	words := []string{"ab", "ba", "z", "zz"}

	anagrams := FindAnagramsInList(words)

	expected := [][]string{[]string{"ab", "ba"}, []string{"z"}, []string{"zz"}}
	if !reflect.DeepEqual(anagrams, expected) {
        t.Errorf("Expected %s, got%s",expected, anagrams)
    }
}
