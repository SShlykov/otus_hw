package linkedlist

// List представляет двусвязный список с произвольным значением типа T
type List[T any] struct {
	length int

	// Head представляет начало списка, предполагается что текущий узел имеет ссылку на предыдущий и следующий узел
	Head *Node[T]
}

// NewList создает новый список с произвольным значением типа T (что-то похожее на конструктор фабрики)
func NewList[T any]() List[T] {
	return List[T]{}
}

// Len возвращает количество элементов в списке
func (ll *List[T]) Len() int {
	return ll.length
}

// Front возвращает первый узел списка
func (ll *List[T]) Front() *Node[T] {
	return ll.Head
}

// Back возвращает последний узел списка
func (ll *List[T]) Back() *Node[T] {
	if ll.Len() == 0 {
		return nil
	}
	last := ll.Head
	for ; last.Next != nil; last = last.Next {
	}
	return last
}

// PushFront добавляет элемент в начало списка
func (ll *List[T]) PushFront(value T) *Node[T] {
	if ll.Len() == 0 {
		newNode := NewNode(value)
		ll.Head = newNode
		ll.length++
		return newNode
	}

	newNode := NewNode(value).SetNext(ll.Head)
	ll.Front().Prev = newNode

	ll.Head = newNode
	ll.length++
	return newNode
}

// PushBack добавляет элемент в конец списка
func (ll *List[T]) PushBack(value T) *Node[T] {
	newNode := NewNode(value)

	if ll.Len() == 0 {
		ll.Head = newNode
		ll.length++
		return newNode
	}

	back := ll.Back()
	newNode.SetPrev(back)
	back.Next = newNode
	ll.length++

	return newNode
}

// Remove удаляет узел из списка
func (ll *List[T]) Remove(node *Node[T]) {
	ll.length--
	switch {
	case node.Prev == nil && node.Next == nil:
	case node.Prev == nil:
		node.Next.Prev = nil
		ll.Head = node.Next
	case node.Next == nil:
		node.Prev.Next = nil
	default:
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
}

// MoveToFront перемещает узел в начало списка
func (ll *List[T]) MoveToFront(node *Node[T]) {
	ll.Remove(node)
	ll.PushFront(node.Value)
}
