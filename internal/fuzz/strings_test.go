package fuzz

import "testing"

func TestSubstituteRandomCharacter(t *testing.T) {
	t.Run("string length is the same", func(t *testing.T) {
		old := "test"
		new := SubstituteRandomCharacter(old)
		if len(new) != len(old) {
			t.Errorf("string length should be %v but is %v\n", len(old), len(new))
		}
	})

	t.Run("string is not equal", func(t *testing.T) {
		old := "test"
		new := SubstituteRandomCharacter(old)
		if new == old {
			t.Errorf("string should not be %v but is %v\n", old, new)
		}
	})
}

func TestRandomCharacterAddition(t *testing.T) {
	t.Run("string length is input + 1", func(t *testing.T) {
		old := "test"
		new := AddRandomCharacter(old)
		if len(new) != (len(old) + 1) {
			t.Errorf("string length should be %v but is %v\n", (len(old) + 1), len(new))
		}
	})
}

func TestRandomCharacterDeletion(t *testing.T) {
	t.Run("string length is input - 1", func(t *testing.T) {
		old := "test"
		new := DeleteRandomCharacter(old)
		if len(new) != (len(old) - 1) {
			t.Errorf("string length should be %v but is %v\n", (len(old) - 1), len(new))
		}
	})
}
