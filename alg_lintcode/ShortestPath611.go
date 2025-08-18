package main

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
	return -1
}
