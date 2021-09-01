/*
@Time   : 2021/8/25 16:49
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
}
