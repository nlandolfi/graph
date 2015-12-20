package graph

import (
	"container/list"
	"fmt"
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

// --- List Queue {{{

// notes: first elem is next

type listQueue []*list.List

func (lq *listQueue) pushSlice(l []*list.List) {
	newqueue := make([]*list.List, len(*lq)+len(l))
	for i := range *lq {
		newqueue[i] = (*lq)[i]
	}
	o := len(*lq)
	for i := range l {
		newqueue[o+i] = l[i]
	}
	*lq = newqueue
}

func (lq *listQueue) push(l *list.List) {
	*lq = append(*lq, l)
}

func (lq *listQueue) pop() *list.List {
	t := (*lq)[0]
	*lq = (*lq)[1:len(*lq)]
	return t
}

func (lq *listQueue) length() int {
	return len(*lq)
}

// --- }}}

// Returns the path from start until a goal (node satisfying 'satisfaction') using the depth first search
func DepthFirstSearch(start *node, satisfaction func(*node) bool) (*list.List, error) {
	// The set of nodes we have already examined, prevents cycles
	seen := make(map[*node]bool)

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
		if _, ok := seen[current]; ok {
			continue
		} else {
			seen[current] = true // else mark it as seen
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

// Returns the path from start until a goal (node satisfying 'satisfaction') using the breadth first search
func BreadthFirstSearch(start *node, satisfaction func(*node) bool) (*list.List, error) {
	// The set of nodes we have already examined, prevents cycles
	seen := make(map[*node]bool)

	// The queue of nodes to examine
	queue := make(listQueue, 1)
	l := list.New()
	l.PushBack(start)
	queue[0] = l

	// The edge tree
	edgeTree := make(map[*node]*node)
	edgeTree[start] = nil

	for queue.length() != 0 {
		currentList := queue.pop()
		current := currentList.Back().Value.(*node)

		// If we have already seen this node, bail
		if _, ok := seen[current]; ok {
			continue
		} else {
			seen[current] = true // else mark it as seen
		}

		// Does this node satisfy (terminate) our search
		if satisfaction(current) {
			return currentList, nil // we can stop now
		}

		for _, v := range *current.edges {
			l := list.New()
			l.PushBackList(currentList)
			l.PushBack(v)
			queue.push(l)
		}
	}

	return list.New(), fmt.Errorf("path not found to satisfaction")
}
