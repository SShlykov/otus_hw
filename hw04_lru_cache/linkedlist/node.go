package linkedlist

// Node представляет узел связного списка с произвольным значением типа T
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// NewNode создает новый узел связного списка
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{Value: value, Next: nil, Prev: nil}
}

// SetNext устанавливает следующий узел,
// возвращает текущий узел
func (n *Node[T]) SetNext(next *Node[T]) *Node[T] {
	n.Next = next
	return n
}

// SetPrev устанавливает предыдущий узел,
// возвращает текущий узел
func (n *Node[T]) SetPrev(next *Node[T]) *Node[T] {
	n.Prev = next
	return n
}
