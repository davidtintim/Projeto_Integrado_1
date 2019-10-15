package foodchain

import "strings"

// TestVersion is the unit test version that this program is built to pass.
const TestVersion = 1

/*Verse sings a verse of "I know an old lady who swallowed a fly".*/
func Verse(v int) string {
	verse := "I know an old lady who swallowed " + first[v]
	if 8 <= v {
		return verse
	}
	for i := v; 0 <= i; i-- {
		verse += refrain[i]
	}
	return verse
}

/*Verses sings a set of verses of "I know an old lady who swallowed a fly".*/
func Verses(start, stop int) string {
	var verses = []string{}
	for v := start; v <= stop; v++ {
		verses = append(verses, Verse(v))
	}
	return strings.Join(verses, "\n\n")
}

/*Song sings "I know an old lady who swallowed a fly" in it's entirty.*/
func Song() string {
	return Verses(1, 8)
}