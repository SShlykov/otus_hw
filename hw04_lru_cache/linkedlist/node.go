package linkedlist

// Node представляет узел связного списка с произвольным значением типа T
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}
