/*
@Time   : 2021/12/30 12:05
@Author : ckx0709
@Remark : 浮点数工具包
*/
package util

import (
	"fmt"
	"math"
	"strconv"
)

/*
浮点数向上/下保留小数点位数，
params float 类型必须位float32/64；否则返回零值
params isUp : 1:向上保留，其他：向下保留
*/
func FloatPositionRetain(float interface{}, point int, isUp int) float64 {
	f, ok := float.(float64)
	if !ok {
		return 0.0
	}
	switch isUp {
	case 1:
		f += math.Pow(0.1, float64(point+1)) * 4.1
	default:
		f -= math.Pow(0.1, float64(point+1)) * 5
	}
	stf := fmt.Sprintf(fmt.Sprintf("%%.%df", point), f)
	parseFloat, _ := strconv.ParseFloat(stf, point)
	return parseFloat
}
