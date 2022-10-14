/*
@Time   : 2022/9/14 15:44
@Author : ckx0709
@Remark :
*/
package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestMatchPoses_Len(t *testing.T) {

	var mts matchPoses
	mts = append(mts, []matchPos{
		{
			distance:  1.1,
			areaIndex: 8,
			posId:     "D",
		},
		{
			distance:  0.1,
			areaIndex: 77,
			posId:     "A",
		},
		{
			distance:  1.3,
			areaIndex: 3,
			posId:     "Q",
		},
	}...)

	sort.Sort(mts)
	fmt.Println(mts)
}

func TestMatchPoses_Less(t *testing.T) {
	mts := matchPosesSortAreaIndex{}
	mts.matchPoses = append(mts.matchPoses, []matchPos{
		{
			distance:  1.1,
			areaIndex: 8,
			posId:     "D",
		},
		{
			distance:  0.1,
			areaIndex: 77,
			posId:     "A",
		},
		{
			distance:  1.3,
			areaIndex: 3,
			posId:     "Q",
		},
	}...)

	sort.Sort(mts)
	fmt.Println(mts, "\t", mts.Len())

}

func TestMatchPoses_Swap(t *testing.T) {

}
