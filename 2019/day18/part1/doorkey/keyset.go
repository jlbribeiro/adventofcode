package doorkey

import "strings"

type KeySet int

func (ks KeySet) HasKey(keyID int) bool {
	bm := keyIDToBitMask(keyID)
	return int(ks)&bm > 0
}

func (ks KeySet) IsSubsetOf(superset KeySet) bool {
	return ks&superset == ks
}

func (ks KeySet) WithKey(key rune) KeySet {
	return ks | KeySet(keyToBitMask(key))
}

func (ks KeySet) WithoutKey(key rune) KeySet {
	return ks & KeySet(^keyToBitMask(key))
}

func (ks KeySet) WithoutKeySet(o KeySet) KeySet {
	return ks & KeySet(^o)
}

func (ks KeySet) String() string {
	if ks == 0 {
		return "none"
	}

	var keys []string
	for keyID := 0; keyIDToBitMask(keyID) <= int(ks); keyID++ {
		if ks.HasKey(keyID) {
			keys = append(keys, string(rune('a'+keyID)))
		}
	}

	return strings.Join(keys, ",")
}
