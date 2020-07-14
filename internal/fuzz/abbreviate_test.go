package fuzz

import "testing"

func TestAbbreviation(t *testing.T) {
	t.Run("requested string length is produced", func(t *testing.T) {
		old := "the quick brown fox"
		new := AbbreviateString(old, 5)
		if len(new) != 5 {
			t.Errorf("string length should be %v but is %v. string: %v", 5, len(new), new)
		}
	})
	t.Run("does not abbreviate if under threshold", func(t *testing.T) {
		old := "t"
		new := AbbreviateString(old, 5)
		if len(new) != 1 {
			t.Errorf("string length should be %v but is %v. string: %v", 1, len(new), new)
		}
	})
}
