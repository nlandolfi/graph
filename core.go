package graph

// Element should really only take pointer types,
// it doesn't make sense to copy values all over
// the place , but then we may need a comparison func
type Element interface{}

type Interface interface {
	Nodes() <-chan Node
}

type Node interface {
	Element() Element
	Edges() <-chan Node
}

/// -- forget about ^^ for now

type node struct {
	//`value  Element
	id    int
	edges *[]*node
}

type graph []*node

func NewNode(id int, edges ...*node) *node {
	return &node{id: id, edges: &edges}
}

func (n *node) SetEdges(edges []*node) {
	n.edges = &edges
}
