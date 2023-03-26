// Package graph. graph is directed graph with weights.
package graph

import (
	"errors"
)

const (
	EdgeAlreadyExists = "attempt to add an already existing edge"
	EdgeNotExists     = "edge with such value do not exist"
	PathNotExists     = "there is no path from the given node to the target node"
)

type graph struct {
	listOfEdges map[string][]*Edge
}

func New() *graph {
	return &graph{listOfEdges: make(map[string][]*Edge)}
}

type Edge struct {
	start  string
	end    string
	weight int
}

func CreateEdge(startNode, endNode string, weight int) *Edge {
	return &Edge{start: startNode, end: endNode, weight: weight}
}

func (g *graph) AddEdge(startNode, endNode string, weight int) {
	edge := CreateEdge(startNode, endNode, weight)
	g.listOfEdges[startNode] = append(g.listOfEdges[startNode], edge)
}

func (g *graph) AddEdges(edges []*Edge) error {
	for _, edge := range edges {
		if g.IsExist(edge) {
			return errors.New(EdgeAlreadyExists)
		}
		g.listOfEdges[edge.start] = append(g.listOfEdges[edge.start], edge)
	}
	return nil
}

func (g *graph) IsExist(edge *Edge) bool {
	list := g.listOfEdges[edge.start]
	for index := range list {
		if list[index].end == edge.end && list[index].weight == edge.weight {
			return true
		}
	}
	return false
}

func (g *graph) ChangeWeightInEdge(startNode, endNode string, newWeight int) error {
	for _, edge := range g.listOfEdges[startNode] {
		if edge.end == endNode {
			edge.weight = newWeight
			return nil
		}
	}
	return errors.New(EdgeNotExists)
}

func (g *graph) RemoveEdgeBetweenNodes(startNode, endNode string) bool {
	list := g.listOfEdges[startNode]
	for i := range list {
		if list[i].end == endNode {
			list[i] = list[len(list)-1]
			g.listOfEdges[startNode] = g.listOfEdges[startNode][:len(list)-1]
			if len(g.listOfEdges[startNode]) == 0 {
				delete(g.listOfEdges, startNode)
			}
			return true
		}
	}
	return false
}

func (g *graph) RemoveNode(node string) bool {
	wasDeletion := false
	for _, startNode := range g.listOfEdges {
		for _, str := range startNode {
			if str.start == node {
				delete(g.listOfEdges, str.start)
				wasDeletion = true
			}
			if str.end == node {
				g.RemoveEdgeBetweenNodes(str.start, str.end)
				wasDeletion = true
			}
		}
	}
	return wasDeletion
}

func (g *graph) Clear() {
	for key, _ := range g.listOfEdges {
		delete(g.listOfEdges, key)
	}
}

// FindShortestPathBFS using BFS algorithm and finds the minimum number of edges between two nodes
func (g *graph) FindShortestPathBFS(initNode, targetNode string) (int, error) {
	if len(g.listOfEdges[initNode]) == 0 {
		return 0, errors.New(PathNotExists)
	}

	if initNode == targetNode {
		return 0, nil
	}

	q := newQueue()
	q.push(initNode)
	alreadyChecked := make(map[string]int)
	minAmountOfEdges := 0

	for q.size != 0 {
		currentNode := q.pop()

		for _, item := range g.listOfEdges[currentNode] {
			q.push(item.end)
			amount := alreadyChecked[item.start] + 1
			alreadyChecked[item.end] = amount
		}

		if currentNode == targetNode {
			minAmountOfEdges = alreadyChecked[currentNode]
			break
		}

	}
	return minAmountOfEdges, nil
}
