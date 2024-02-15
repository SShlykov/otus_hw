package hw03frequencyanalysis

import (
	"strings"
	"unicode"
)

// Top10 возвращает 10 самых часто встречающихся слов в порядке убывания частоты.
// Если слова с одинаковой частотой, то сортируются лексикографически.
func Top10(text string) []string {
	wordSet := makeWordSet(text)
	wordHeap := NewWordHeap()
	for word, freq := range wordSet {
		wordHeap.Add(word, freq)
	}

	return wordHeap.TopN(10)
}

// makeWordSet создает не стандартное множество слов из текста (вместо bool хранит частоту слова)
func makeWordSet(text string) map[string]int {
	wordSet := make(map[string]int)

	words := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-'
	})

	for _, word := range words {
		word = strings.ToLower(word)

		if word == "" || word == "-" {
			continue
		}
		wordSet[word]++
	}

	return wordSet
}
