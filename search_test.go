package graph

import (
	"container/list"
	"testing"
)

var llist *list.List

func BenchmarkDFS(b *testing.B) {
	// run the Fib function b.N times
	S := NewNode(1)
	A := NewNode(2)
	B := NewNode(3)
	C := NewNode(4)
	G := NewNode(5)

	S.edges = &[]*node{A, B} // note b wins tiebreak
	A.edges = &[]*node{B, C} // note c wins tiebreak
	B.edges = &[]*node{C}
	C.edges = &[]*node{G}

	var l *list.List

	for n := 0; n < b.N; n++ {
		l, _ = DepthFirstSearch(S, func(n *node) bool {
			return n == G
		})
	}

	llist = l
}

func BenchmarkBFS(b *testing.B) {
	// run the Fib function b.N times
	S := NewNode(1)
	A := NewNode(2)
	B := NewNode(3)
	C := NewNode(4)
	G := NewNode(5)

	S.edges = &[]*node{A, B} // note b wins tiebreak
	A.edges = &[]*node{B, C} // note c wins tiebreak
	B.edges = &[]*node{C}
	C.edges = &[]*node{G}

	var l *list.List

	for n := 0; n < b.N; n++ {
		l, _ = BreadthFirstSearch(S, func(n *node) bool {
			return n == G
		})
	}

	llist = l
}
