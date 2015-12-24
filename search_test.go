package graph

import (
	"container/list"
	"testing"
)

/*
               _______A_______             1
			 ↙        ↓       ↘
		   _B_       _C_       _D_         3
     	 ↙  ↓ ↘    ↙  ↓ ↘    ↙  ↓ ↘
		 E  F  G  H   I  J  K   L  M       9

		 DFS: Á → B → E → F → G → C → H → I → J → D → K → L → M
		 BFS: A → B → C → D → E → F → G → H → I → J → K → L → M
*/

func TestDFS(t *testing.T) {
	nodes := make([]*node, 13)
	c := 0

	A := NewNode(1)
	B := NewNode(2)
	C := NewNode(3)
	D := NewNode(4)
	E := NewNode(5)
	F := NewNode(6)
	G := NewNode(7)
	H := NewNode(8)
	I := NewNode(9)
	J := NewNode(10)
	K := NewNode(11)
	L := NewNode(12)
	M := NewNode(13)

	nodeMap := map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
	}

	A.edges = &[]*node{D, C, B}
	B.edges = &[]*node{G, F, E}
	C.edges = &[]*node{J, I, H}
	D.edges = &[]*node{M, L, K}

	DepthFirstSearch(A, func(n *node) bool {
		nodes[c] = n
		c++
		return false // continue search
	})

	if c != 13 {
		t.Fatalf("Expected to see 13 nodes on depth search")
	}

	// DFS: Á → B → E → F → G → C → H → I → J → D → K → L → M
	expectedOrder := []string{"A", "B", "E", "F", "G", "C", "H", "I", "J", "D", "K", "L", "M"}

	for i := range nodes {
		if nodeMap[nodes[i].id] != expectedOrder[i] {
			t.Fatalf("Expected node %s, but got %s in index %d", expectedOrder[i], nodeMap[nodes[i].id], i)
		}
	}

	c = 0
	BreadthFirstSearch(A, func(n *node) bool {
		nodes[c] = n
		c++
		return false // continue search
	})

	if c != 13 {
		t.Fatalf("Expected to see 13 nodes on breadth first search")
	}

	// BFS: A → B → C → D → E → F → G → H → I → J → K → L → M
	expectedOrder = []string{"A", "D", "C", "B", "M", "L", "K", "J", "I", "H", "G", "F", "E"}

	for i := range nodes {
		if nodeMap[nodes[i].id] != expectedOrder[i] {
			t.Fatalf("Expected node %s, but got %s in index %d", expectedOrder[i], nodeMap[nodes[i].id], i)
		}
	}
}

// --- Benchmarks {{{

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

// --- }}}
