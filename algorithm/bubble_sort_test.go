/*
@Time   : 2021/12/24 16:57
@Author : ckx0709
@Remark :
*/
package algorithm

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nubs := []int{5, 3, 2, 1, 9, 4, 7, 5, 8, 6, 3}
	sort := BubbleSort(nubs)
	fmt.Printf("%+v", sort)
}
