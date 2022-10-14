// Package sort
/*
@Time   : 2022/9/14 15:41
@Author : ckx0709
@Remark :
*/
package sort

type matchPos struct {
	distance  float64
	areaIndex int
	posId     string
}

//结构体切片 需要实现sort.interface
type matchPoses []matchPos

func (matches matchPoses) Len() int {
	return len(matches)
}

func (matches matchPoses) Less(i, j int) bool {
	/*排序的依据field 以及 升降序控制
	< : 降序
	> : 升序
	*/
	return matches[i].distance > matches[j].distance
}

func (matches matchPoses) Swap(i, j int) {
	matches[i], matches[j] = matches[j], matches[i]
}

//以继承的方式 重写Less方法实现不同field排序方式
type matchPosesSortAreaIndex struct {
	matchPoses
}

func (m matchPosesSortAreaIndex) Less(i, j int) bool {
	return m.matchPoses[i].posId < m.matchPoses[j].posId
}
