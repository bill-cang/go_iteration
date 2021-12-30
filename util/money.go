/*
@Time   : 2021/12/30 9:43
@Author : ckx0709
@Remark :
*/
package util

//涨停板->倍率
func GetTimes(limitBoards int) (times float64) {
	var a float64 = 1
	limitBoards++
	for i := 0; i < limitBoards; i++ {
		b := a * 0.1
		a += b
	}
	return a / 1
}
