/*
@Time   : 2021/12/24 16:52
@Author : ckx0709
@Remark : 冒泡排序
*/
package algorithm

//
func BubbleSort(numbers []int) (sortNub []int) {

	//节约内存分配次数
	temp := 0
	for i := 0; i < len(numbers); i++ {
		for k := i + 1; k < len(numbers); k++ {
			if numbers[i] > numbers[k] {
				temp = numbers[i]
				numbers[i] = numbers[k]
				numbers[k] = temp
			}
		}
	}

	return numbers
}
