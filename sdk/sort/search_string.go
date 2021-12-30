/*
@Time   : 2021/12/29 15:41
@Author : ckx0709
@Remark :
*/
package sort

import "sort"

func SearchStringFromSplice(strSpl []string, str string) int {
	/*
	SearchStrings在递增顺序的a中搜索x，返回x的索引。
	如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	*/
	return sort.SearchStrings(strSpl, str)
}


