package linkedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(1)

	require.Equal(t, 1, node.Value)
	require.Nil(t, node.Next)
	require.Nil(t, node.Prev)
}

func TestNode_SetNext(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)

	node1.SetNext(node2)

	require.Equal(t, 2, node1.Next.Value)
	require.Nil(t, node2.Next)
}

func TestNode_SetPrev(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)

	node1.SetPrev(node2)

	require.Equal(t, 2, node1.Prev.Value)
	require.Nil(t, node2.Prev)
}
