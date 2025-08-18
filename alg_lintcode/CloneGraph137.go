package main

import (
	"container/list"
)

/**
 */
type UndirectedGraphNode struct {
	Label     int
	Neighbors []*UndirectedGraphNode
}

/**
 * @param node: A undirected graph node
 * @return: A undirected graph node
 */
func CloneGraph(node *UndirectedGraphNode) *UndirectedGraphNode {
	// write your code here
	if node == nil {
		return node
	}
	// 首先进行克隆，形成新老映射
	mapping := clone(node)
	//fmt.Println(mapping)
	//建联
	listFound(mapping)

	return mapping[node]
}

func listFound(mapping map[*UndirectedGraphNode]*UndirectedGraphNode) {
	for k, v := range mapping {
		neighbors := k.Neighbors
		for _, nei := range neighbors {
			node := mapping[nei]
			v.Neighbors = append(v.Neighbors, node)
		}
	}
}

func clone(node *UndirectedGraphNode) map[*UndirectedGraphNode]*UndirectedGraphNode {
	mapping := make(map[*UndirectedGraphNode]*UndirectedGraphNode)
	visited := make(map[*UndirectedGraphNode]interface{})
	queue := list.New()
	visited[node] = true
	queue.PushBack(node)
	for queue.Len() > 0 {
		front := queue.Front()
		oldNode := front.Value.(*UndirectedGraphNode)
		queue.Remove(front)
		var newNode = &UndirectedGraphNode{
			Label:     oldNode.Label,
			Neighbors: make([]*UndirectedGraphNode, 0),
		}
		mapping[oldNode] = newNode
		for _, nodeValue := range oldNode.Neighbors {
			if _, exist := visited[nodeValue]; exist {
				continue
			}
			queue.PushBack(nodeValue)
			visited[nodeValue] = true
		}
	}
	return mapping
}
