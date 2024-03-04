package hw03frequencyanalysis

import (
	"container/heap"
	"testing"
)

func TestWordHeap_TopN(t *testing.T) {
	h := NewWordHeap()
	words := []WordFrequency{
		{"word", 2},
		{"test", 3},
		{"another", 1},
	}
	for _, wf := range words {
		h.Add(wf.Word, wf.Frequency)
	}
	top2 := h.TopN(2)
	expected := []string{"test", "word"}
	for i, word := range top2 {
		if word != expected[i] {
			t.Errorf("Expected %s, got %s at position %d", expected[i], word, i)
		}
	}
	if h.Len() != 1 {
		t.Errorf("Heap size incorrect after TopN, expected 1, got %d", h.Len())
	}
}

func TestWordHeap_TopN_Overflow(t *testing.T) {
	h := NewWordHeap()
	words := []WordFrequency{
		{"word", 2},
		{"test", 3},
		{"another", 1},
	}
	for _, wf := range words {
		h.Add(wf.Word, wf.Frequency)
	}
	top10 := h.TopN(10)
	expected := []string{"test", "word", "another"}
	for i, word := range top10 {
		if word != expected[i] {
			t.Errorf("Expected %s, got %s at position %d", expected[i], word, i)
		}
	}
	if h.Len() != 0 {
		t.Errorf("Heap size incorrect after TopN, expected 1, got %d", h.Len())
	}
}

func TestWordHeap_Order(t *testing.T) {
	h := NewWordHeap()
	words := []WordFrequency{
		{"word", 2},
		{"test", 3},
		{"another", 1},
	}
	for _, wf := range words {
		heap.Push(h, wf)
	}
	if (*h)[0].Word != "test" {
		t.Errorf("Top element should be 'test', got '%s'", (*h)[0].Word)
	}
}

func TestWordHeap_Add(t *testing.T) {
	h := NewWordHeap()
	h.Add("test", 1)
	h.Add("word", 2)
	if h.Len() != 2 {
		t.Errorf("Heap size incorrect, expected 2, got %d", h.Len())
	}
}
