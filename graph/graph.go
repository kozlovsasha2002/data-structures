// Package graph. graph is directed graph with weights.
package graph

import (
	"errors"
)

const (
	NodeAlreadyExists = "a node with this value already exists"
	NodeNotExists     = "one of the nodes does not exist"
	EdgeNotExists     = "edges with such nodes do not exist"
)

type graph struct {
	listOfNodes map[interface{}]*node
}

func New() *graph {
	return &graph{listOfNodes: make(map[interface{}]*node)}
}

type Edge struct {
	dataOfStartNode string
	dataOfEndNode   string
	weight          int
}

func CreateEdge(dataOfStartNode, dataOfEndNode string, weight int) *Edge {
	return &Edge{dataOfStartNode: dataOfStartNode, dataOfEndNode: dataOfEndNode, weight: weight}
}

type node struct {
	data      string
	neighbors map[*node]int
}

func createNode(data string) *node {
	return &node{data: data, neighbors: make(map[*node]int)}
}

func (g *graph) AddNode(data string) error {
	newNode := createNode(data)

	if g.isExist(newNode) {
		return errors.New(NodeAlreadyExists)
	}

	g.listOfNodes[data] = newNode
	return nil
}

func (g *graph) isExist(n *node) bool {
	if _, ok := g.listOfNodes[n.data]; ok {
		return true
	}
	return false
}

func (g *graph) AddAllNodes(data []string) error {
	for _, item := range data {
		err := g.AddNode(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *graph) AddEdge(start, end string, weight int) {
	if !g.isExist(createNode(start)) {
		g.AddNode(start)
	}

	if !g.isExist(createNode(end)) {
		g.AddNode(end)
	}

	startNode, _ := g.listOfNodes[start]
	endNode, _ := g.listOfNodes[end]
	startNode.neighbors[endNode] = weight
}

func (g *graph) AddEdges(edges []*Edge) {
	for _, edge := range edges {
		g.AddEdge(edge.dataOfStartNode, edge.dataOfEndNode, edge.weight)
	}
}

func (g *graph) RemoveEdgeBetweenNodes(start, end string) error {
	if !g.isExist(createNode(start)) || !g.isExist(createNode(end)) {
		return errors.New(NodeNotExists)
	}

	startNode, _ := g.listOfNodes[start]
	if _, ok := startNode.neighbors[createNode(end)]; ok {

	}
	return nil
}

//удалить вершину, при удалении вершины удаляются все ребра с ней взаимосвязанные

//изменить вес в ребре, если ребра не существует ошибка

//найти кратчайший путь по алгоритмы Дейкстры (он не работает с отрицательными ребрами

//найти кратчайший путь поиском в ширину (в данном алгоритме не учитываются веса ребёр)

//найти кратчайший путь по алгоритму Форда-Баллмана, который работает с отрицательными ребрами
