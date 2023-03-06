package main

import (
	"fmt"
	"math"
	"unsafe"

	"dijkstra-shortest-path/queue"

	g "github.com/zyedidia/generic"
	"github.com/zyedidia/generic/hashset"
)

// TODO: watch https://www.youtube.com/watch?v=BuvKtCh0SKk

type Node struct {
	// node name/id
	name string
	// distance from source node
	distance int
	// list of shortest path from source node until this node (exclusive: this node is not encluded)
	shortestPath []*Node
	// adjacent nodes mapped to their edge weight
	adjacentNodes map[*Node]int
}

func NewNode(name string) *Node {
	return &Node{
		name:          name,
		distance:      math.MaxInt,
		shortestPath:  make([]*Node, 0, 100),
		adjacentNodes: make(map[*Node]int),
	}
}

func (n *Node) AddAdjacent(node *Node, weight int) {
	n.adjacentNodes[node] = weight
}

func CalculateShortestPath(source *Node) {
	// source node distance from source is zero
	source.distance = 0
	// settled nodes are nodes that have their all adjacent nodes' min distance evaluated
	settledNodes := hashset.New(
		/*capacity: */ 100,
		/*equals: */ g.Equals[*Node],
		/*hash: */
		func(n *Node) uint64 { return g.HashUint64(uint64(uintptr(unsafe.Pointer(n)))) },
	)
	// unsettled nodes are the adjacent nodes that their min distance is not evaluated yet
	unsettledNodes := queue.NewMinPriority(
		/*priorityValue: */ func(n *Node) int { return n.distance },
	)
	unsettledNodes.Enqueue(source)

	for !unsettledNodes.Empty() {

		currentNode := unsettledNodes.Dequeue()

		for adjacentNode, weight := range currentNode.adjacentNodes {
			if !settledNodes.Has(adjacentNode) {
				evaluateDistanceAndPath(adjacentNode, weight, currentNode)
				unsettledNodes.Enqueue(adjacentNode)
			}
		}

		settledNodes.Put(currentNode)
	}
}

func evaluateDistanceAndPath(
	adjacentNode *Node,
	edgeWeight int,
	sourceNode *Node,
) {
	newDistance := sourceNode.distance + edgeWeight
	if newDistance < adjacentNode.distance {
		adjacentNode.distance = newDistance
		adjacentNode.shortestPath = append(sourceNode.shortestPath, sourceNode)
	}
}

func PrintShortestPath(node *Node) {
	var s string
	for _, n := range node.shortestPath {
		s += fmt.Sprintf("%s -> ", n.name)
	}
	s += fmt.Sprintf("%s: %d", node.name, node.distance)
	fmt.Println(s)
}

// -----------------------------------------------------------------------------

func main() {
	A := NewNode("A")
	B := NewNode("B")
	C := NewNode("C")
	D := NewNode("D")
	E := NewNode("E")
	F := NewNode("F")

	A.AddAdjacent(B, 2)
	A.AddAdjacent(C, 4)

	B.AddAdjacent(C, 3)
	B.AddAdjacent(D, 1)
	B.AddAdjacent(E, 5)

	C.AddAdjacent(D, 2)

	D.AddAdjacent(E, 1)
	D.AddAdjacent(F, 4)

	E.AddAdjacent(F, 2)

	// calculate shortest path from A
	CalculateShortestPath(A)

	PrintShortestPath(A)
	PrintShortestPath(B)
	PrintShortestPath(C)
	PrintShortestPath(D)
	PrintShortestPath(E)
	PrintShortestPath(F)
}
