package main

import "testing"

func TestRandomWordSubstitution(t *testing.T) {
	dictionary = []string{"test", "word"}
	t.Run("word length is the same on input as output", func(t *testing.T) {
		old := "the quick brown fox"
		new := substituteWordRandom(old)
		if len(getWords(old)) != len(getWords(old)) {
			t.Errorf("word amount should be %v but is %v\n", len(getWords(old)), len(getWords(new)))
		}
	})
}

func TestRandomWordAddition(t *testing.T) {
	t.Run("word length is input + 1", func(t *testing.T) {
		dictionary = []string{"test", "word"}
		old := "the quick brown fox"
		new := addWordRandom(old)
		if len(getWords(new)) != (len(getWords(old)) + 1) {
			t.Errorf("word amount should be %v but is %v\n", (len(getWords(old)) + 1), len(getWords(new)))
		}
	})
}

func TestRandomWordDeletion(t *testing.T) {
	t.Run("word length is input - 1", func(t *testing.T) {
		old := "the quick brown fox"
		new := deleteWordRandom(old)
		if len(getWords(new)) != (len(getWords(old)) - 1) {
			t.Errorf("word amount should be %v but is %v\n", (len(getWords(old)) - 1), len(getWords(new)))
		}
	})
}
