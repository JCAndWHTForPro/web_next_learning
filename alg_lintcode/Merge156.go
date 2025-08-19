package main

import "sort"

/**
 * Definition of Interval:
 *
 */
type Interval struct {
	Start, End int
}

/*
*
描述
我们以一个 Interval 类型的列表 intervals 来表示若干个区间的集合，其中单个区间为 (start, end)。
你需要合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

你需要对有重叠部分的区间进行合并
即使某个区间与其他区间都没有重叠，也需要将其输出
*/
func Merge(intervals []*Interval) []*Interval {
	// write your code here
	l := len(intervals)
	if l == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i := 1; i < len(intervals); {
		curInterval := intervals[i]
		preInterval := intervals[i-1]
		if curInterval.Start > preInterval.End {
			i++
			continue
		}
		if curInterval.End > preInterval.End {
			preInterval.End = curInterval.End
		}
		intervals = append(intervals[:i], intervals[i+1:]...)

	}
	return intervals
}
