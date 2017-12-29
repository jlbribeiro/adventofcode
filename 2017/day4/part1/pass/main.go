package pass

import "strings"

// ValidPassphrase checks whether a passphrase contains repeated words.
func ValidPassphrase(passphrase string) bool {
	words := strings.Split(passphrase, " ")

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i] == words[j] {
				return false
			}
		}
	}

	return true
}
