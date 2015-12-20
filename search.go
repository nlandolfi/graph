package graph

import (
	"container/list"
	"fmt"

	"github.com/nlandolfi/set"
)

// --- Node Stack {{{

type nodeStack []*node

func (ns *nodeStack) pushSlice(n []*node) {
	newstack := make([]*node, len(*ns)+len(n))
	for i := range *ns {
		newstack[i] = (*ns)[i]
	}
	o := len(*ns)
	for i := range n {
		newstack[o+i] = n[i]
	}
	*ns = newstack
}

func (ns *nodeStack) push(n *node) {
	*ns = append(*ns, n)
}

func (ns *nodeStack) pop() *node {
	t := (*ns)[len(*ns)-1]
	*ns = (*ns)[0 : len(*ns)-1]
	return t
}

func (ns *nodeStack) length() int {
	return len(*ns)
}

// --- }}}

// --- Node Queue {{{

type nodeQueue []*node

func (nq *nodeQueue) pushSlice(n []*node) {
	newqueue := make([]*node, len(*nq)+len(n))
	for i := range *nq {
		newqueue[i] = (*nq)[i]
	}
	o := len(*nq)
	for i := range n {
		newqueue[o+i] = n[i]
	}
}

func (nq *nodeQueue) push(n *node) {
	*nq = append(*nq, n)
}

func (nq *nodeQueue) pop() *node {
	t := (*nq)[0]
	*nq = (*nq)[1:len(*nq)]
	return t
}

func (nq *nodeQueue) length() int {
	return len(*nq)
}

// --- }}}

// Returns the path from start until a goal (node satisfying 'satisfaction') using the depth first search
func DepthFirstSearch(start *node, satisfaction func(*node) bool) (*list.List, error) {
	// The set of nodes we have already examined, prevents cycles
	seen := set.New()

	// The path we are currently examining
	path := list.New()

	// The stack of nodes to examine
	stack := make(nodeStack, 1)
	stack[0] = start

	// The levels of each node, (number of edges away from start)
	levels := make(map[*node]int)
	levels[start] = 0
	level := 0

	for stack.length() != 0 {
		current := stack.pop()

		// If we have already seen this node, bail
		if seen.Contains(current) {
			continue
		} else {
			seen.Add(current) // else mark it as seen
		}

		// Have we dropped down a level?
		if levels[current] < level {
			path.Remove(path.Back())
		}

		// This is now a member of our path
		path.PushBack(current)

		// Does this node satisfy (terminate) our search
		if satisfaction(current) {
			return path, nil // we can stop now
		}

		// otherwise, we want to look at all the slice on the next level
		level += 1
		stack.pushSlice(*current.edges)
		for _, v := range *current.edges {
			levels[v] = level
		}
	}

	return path, fmt.Errorf("path not found to satisfaction")
}
