package graph

import "testing"

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

	for n := 0; n < b.N; n++ {
		DepthFirstSearch(S, func(n *node) bool {
			return n == G
		})
	}
}
