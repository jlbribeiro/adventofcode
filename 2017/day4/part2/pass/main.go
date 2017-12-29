package pass

import (
	"sort"
	"strings"
)

// CanonicalWord returns the canonical form of word; that means it sorts all the
// letters so that comparing canonical words may be used to test if two words
// are anagrams of each other.
func CanonicalWord(word string) string {
	chars := []byte(word)

	sort.Slice(chars, func(a, b int) bool {
		return chars[a] < chars[b]
	})

	return string(chars)
}

// ValidPassphrase checks whether a passphrase contains repeated words.
func ValidPassphrase(passphrase string) bool {
	words := strings.Split(passphrase, " ")

	for i := range words {
		words[i] = CanonicalWord(words[i])
	}

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i] == words[j] {
				return false
			}
		}
	}

	return true
}
