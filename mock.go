package graph

import (
	"fmt"
	"log"
)

func Test() {

	S := NewNode(1)
	A := NewNode(2)
	B := NewNode(3)
	C := NewNode(4)
	G := NewNode(5)

	m := map[int]string{
		1: "S",
		2: "A",
		3: "B",
		4: "C",
		5: "G",
	}

	S.edges = &[]*node{A, B} // note b wins tiebreak
	A.edges = &[]*node{B, C} // note c wins tiebreak
	B.edges = &[]*node{C}
	C.edges = &[]*node{G}

	l, err := DepthFirstSearch(S, func(n *node) bool {
		return n == G
	})

	if err != nil {
		log.Fatal(err)
	}

	e := l.Front()

	s := ""
	for e != nil {
		s += fmt.Sprintf("%s, ", m[e.Value.(*node).id])
		e = e.Next()
	}

	log.Printf("%s", s[0:len(s)-1])

	l, err = DepthFirstSearch(G, func(n *node) bool {
		return n == S
	})

	// if you start with g, there is no path to s
	if err != nil {
		log.Printf("No path from G to S")
	} else {
		log.Printf("The code is wrong, there should be no path from G to S")
	}

	// -- Breadth First Search

	l, err = BreadthFirstSearch(S, func(n *node) bool {
		return n == G
	})

	if err != nil {
		log.Fatal(err)
	}

	e = l.Front()

	s = ""
	for e != nil {
		s += fmt.Sprintf("%s, ", m[e.Value.(*node).id])
		e = e.Next()
	}

	log.Printf("%s", s[0:len(s)-1])
}
