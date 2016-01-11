package graph

// Element should really only take pointer types,
// it doesn't make sense to copy values all over
// the place , but then we may need a comparison func
type Element interface{}

type Interface interface {
	Nodes() <-chan Node
}

type Node interface {
	ID() int
	Edges() []Node
	SetEdges([]Node)
}

/// -- forget about ^^ for now

type node struct {
	//`value  Element
	id    int
	edges []Node
}

func (n *node) ID() int {
	return n.id
}

func (n *node) Edges() []Node {
	nodes := make([]Node, len(n.edges))
	for i, e := range n.edges {
		nodes[i] = e
	}
	return nodes
}

type graph []*node

func NewNode(id int, edges ...Node) *node {
	return &node{id: id, edges: edges}
}

func (n *node) SetEdges(edges []Node) {
	n.edges = edges
}
