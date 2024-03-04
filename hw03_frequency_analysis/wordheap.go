package hw03frequencyanalysis

import "container/heap"

// WordFrequency - структура для хранения слова и его частоты
type WordFrequency struct {
	Word      string
	Frequency int
}

// WordHeap - куча для хранения частот слов; Len, Less, Swap, Push, Pop - методы для реализации интерфейса heap.Interface
type WordHeap []WordFrequency

func NewWordHeap() *WordHeap {
	h := &WordHeap{}
	heap.Init(h)
	return h
}

func (h WordHeap) Len() int { return len(h) }
func (h WordHeap) Less(i, j int) bool {
	if h[i].Frequency == h[j].Frequency {
		return h[i].Word < h[j].Word
	}
	return h[i].Frequency > h[j].Frequency
}
func (h WordHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *WordHeap) Push(x interface{}) {
	*h = append(*h, x.(WordFrequency))
}
func (h *WordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Add добавляет слово с его частотой в кучу
func (h *WordHeap) Add(word string, freq int) {
	heap.Push(h, WordFrequency{word, freq})
}

// TopN возвращает N самых часто встречающихся слов в порядке убывания частоты.
// Куча удаляет элементы с вершины, поэтому после применения куча будет изменена, если необходимо
// сохранить исходное состояние, то нужно создать копию кучи. Это оставим на усмотрение пользователя.
func (h *WordHeap) TopN(n int) []string {
	topWords := make([]string, 0, n)
	for i := 0; i < n && h.Len() > 0; i++ {
		topWords = append(topWords, heap.Pop(h).(WordFrequency).Word)
	}
	return topWords
}
