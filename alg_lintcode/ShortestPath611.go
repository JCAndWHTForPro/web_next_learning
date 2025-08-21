package main

import "container/list"

/*
*描述
给定骑士在棋盘上的 初始 位置(一个2进制矩阵 0 表示空 1 表示有障碍物)，找到到达 终点 的最短路线，返回路线的长度。如果骑士不能到达则返回 -1 。
起点跟终点必定为空.
骑士不能碰到障碍物.
路径长度指骑士走的步数.
如果骑士的位置为 (x,y)，他下一步可以到达以下这些位置:

(x + 1, y + 2)
(x + 1, y - 2)
(x - 1, y + 2)
(x - 1, y - 2)
(x + 2, y + 1)
(x + 2, y - 1)
(x - 2, y + 1)
(x - 2, y - 1)
*/
type Point struct {
	X, Y int
}

func ShortestPath(grid [][]bool, source *Point, destination *Point) int {
	deltaX := []int{1, 1, -1, -1, 2, 2, -2, -2}
	deltaY := []int{2, -2, 2, -2, 1, -1, 1, -1}
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	if source.X == destination.X && source.Y == destination.Y {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	rst := -1
	queue := list.New()
	visited := make(map[int]interface{})
	queue.PushBack(source)
	visited[source.X*m+source.Y] = true
outter:
	for queue.Len() != 0 {
		curLen := queue.Len()
		rst++
		for i := 0; i < curLen; i += 1 {
			front := queue.Front()
			curPoint := front.Value.(*Point)
			queue.Remove(front)
			if curPoint.X == destination.X && curPoint.Y == destination.Y {
				break outter
			}
			for j := 0; j < len(deltaX); j += 1 {
				dlx := deltaX[j]
				dly := deltaY[j]
				nextx := curPoint.X + dlx
				nexty := curPoint.Y + dly
				if nextx < 0 || nextx >= n || nexty < 0 || nexty >= m || grid[nextx][nexty] == true {
					continue
				}
				visit := nextx*m + nexty
				if _, exsit := visited[visit]; exsit {
					continue
				}
				queue.PushBack(&Point{
					X: nextx,
					Y: nexty,
				})
				visited[visit] = true

			}
		}

	}
	if rst == 0 {
		return -1
	}
	return rst
}
