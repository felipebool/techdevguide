package main

import "fmt"

/*
Given a string S and a set of words D, find the longest word in D that
is a subsequence of S.

Word W is a subsequence of S if some number of characters, possibly zero,
can be deleted from S to form W, without reordering the remaining characters.

Note: D can appear in any format (list, hash table, prefix tree, etc.

For example,
	given the input of
		S = "abppplee" and
		D = {"able", "ale", "apple", "bale", "kangaroo"}
	the correct output would be "apple"

The words "able" and "ale" are both subsequences of S,
but they are shorter than "apple".

The word "bale" is not a subsequence of S because even though S has all
the right letters, they are not in the right order.

The word "kangaroo" is the longest word in D, but it isn't a subsequence of S.
 */


type sequence struct {
	word string
	length int
}

func (s sequence) CheckAgainst(against string) (bool, int) {
	againstLength := len(against)
	sequenceCurrentPosition := 0
	found := 0

	if len(against) > s.length {
		return false, 0
	}

	for _, charAgainst := range against {
		for index, charSequence := range s.word[sequenceCurrentPosition:] {
			if charAgainst == charSequence {
				sequenceCurrentPosition = index + 1
				found++
				break
			}
		}
	}

	if found == againstLength {
		return true, againstLength
	}

	return false, 0
}

func main() {
	entrySequence := "abppplee"

	s := sequence{entrySequence, len(entrySequence)}
	result := sequence{}

	d := []string{"able", "ale", "apple", "bale", "kangaroo"}

	for _, value := range d {
		if itIs, size := s.CheckAgainst(value); itIs {
			result.word = value
			result.length = size
		}
	}

	fmt.Println(result)
}

