package main

/*
BenchmarkFindIndex-6            12137677                98.74 ns/op
BenchmarkFindIndex2-6           15831279                75.01 ns/op

解答失败: 测试用例:[-3,4,3,90] 0 测试结果:[] 期望结果:[0,2] stdout:
*/

import (
	"fmt"
	"sort"
)

var (
	/*	numbers = []int{1, 3, 11, 5, 2, 15, 7}
		target  = 9*/
	numbers = []int{-3, 4, 3, 90}
	target  = 0
)

func FindIndex(soc []int, target int) (ids [2]int, bo bool) {
	ln := len(soc)
	ignoreSet := make(map[int]struct{}, ln)

	for i := 0; i < ln; i++ {
		iv := soc[i]
		_, ok := ignoreSet[i]
		if iv > target || ok {
			continue
		}
		for k := i + 1; k < ln; k++ {
			kv := soc[k]
			if kv > target {
				ignoreSet[k] = struct{}{}
				continue
			}

			if iv+kv == target {
				ids = [2]int{i, k}
				bo = true
				return
			}
		}
	}

	return
}

func FindIndex2(soc []int, target int) (ids [2]int, bo bool) {
	sort.Ints(soc)
	ln := len(soc)
	for i := 0; i < ln; i++ {
		iv := soc[i]

		if iv >= target {
			return
		}

		for k := i + 1; k < ln; k++ {
			kv := soc[k]
			if iv+kv == target {
				ids = [2]int{i, k}
				bo = true
				return
			}

			if kv > target {
				break
			}
		}
	}

	return
}

func twoSum(nums []int, target int) []int {
	ln := len(nums)

	for i := 0; i < ln; i++ {
		iv := nums[i]
		for k := i + 1; k < ln; k++ {
			kv := nums[k]
			if iv+kv == target {
				return []int{i, k}
			}
		}
	}

	return nil
}

func main() {
	//ids, _ := FindIndex(numbers, target)
	//ids, _ := FindIndex2(numbers, target)
	ids := twoSum(numbers, target)
	fmt.Println(ids)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
