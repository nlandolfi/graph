package graph_test

import (
	"container/list"
	"testing"

	"github.com/nlandolfi/graph"
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
	nodes := make([]graph.Node, 13)
	c := 0

	A := graph.NewNode(1)
	B := graph.NewNode(2)
	C := graph.NewNode(3)
	D := graph.NewNode(4)
	E := graph.NewNode(5)
	F := graph.NewNode(6)
	G := graph.NewNode(7)
	H := graph.NewNode(8)
	I := graph.NewNode(9)
	J := graph.NewNode(10)
	K := graph.NewNode(11)
	L := graph.NewNode(12)
	M := graph.NewNode(13)

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

	A.SetEdges([]graph.Node{D, C, B})
	B.SetEdges([]graph.Node{G, F, E})
	C.SetEdges([]graph.Node{J, I, H})
	D.SetEdges([]graph.Node{M, L, K})

	graph.DepthFirstSearch(A, func(n graph.Node) bool {
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
		if nodeMap[nodes[i].ID()] != expectedOrder[i] {
			t.Fatalf("Expected node %s, but got %s in index %d", expectedOrder[i], nodeMap[nodes[i].ID()], i)
		}
	}

	c = 0
	graph.BreadthFirstSearch(A, func(n graph.Node) bool {
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
		if nodeMap[nodes[i].ID()] != expectedOrder[i] {
			t.Fatalf("Expected node %s, but got %s in index %d", expectedOrder[i], nodeMap[nodes[i].ID()], i)
		}
	}
}

// --- Benchmarks {{{

var llist *list.List

func BenchmarkDFS(b *testing.B) {
	// run the Fib function b.N times
	S := graph.NewNode(1)
	A := graph.NewNode(2)
	B := graph.NewNode(3)
	C := graph.NewNode(4)
	G := graph.NewNode(5)

	S.SetEdges([]graph.Node{A, B}) // note b wins tiebreak
	A.SetEdges([]graph.Node{B, C}) // note c wins tiebreak
	B.SetEdges([]graph.Node{C})
	C.SetEdges([]graph.Node{G})

	var l *list.List

	for n := 0; n < b.N; n++ {
		l, _ = graph.DepthFirstSearch(S, func(n graph.Node) bool {
			return n.ID() == G.ID()
		})
	}

	llist = l
}

func BenchmarkBFS(b *testing.B) {
	// run the Fib function b.N times
	S := graph.NewNode(1)
	A := graph.NewNode(2)
	B := graph.NewNode(3)
	C := graph.NewNode(4)
	G := graph.NewNode(5)

	S.SetEdges([]graph.Node{A, B}) // note b wins tiebreak
	A.SetEdges([]graph.Node{B, C}) // note c wins tiebreak
	B.SetEdges([]graph.Node{C})
	C.SetEdges([]graph.Node{G})

	var l *list.List

	for n := 0; n < b.N; n++ {
		l, _ = graph.BreadthFirstSearch(S, func(n graph.Node) bool {
			return n.ID() == G.ID()
		})
	}

	llist = l
}

// --- }}}
