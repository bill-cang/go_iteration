/*
@Time   : 2021/12/17 9:43
@Author : ckx0709
@Remark :
*/
package util

import (
	"fmt"
	"testing"
)

func TestConv2DEC3(t *testing.T) {

	//2021年 被招联金融忽悠借贷5W的实际年利率
	var total float64 = 50000
	var supplies = []float64{4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67, 4166.67}
	var rate = []float64{187.5, 187.5, 187.5, 187.5, 187.5, 187.5, 187.5, 187.5, 187.5, 187.5, 187.5, 187.5}

	annualized, a := GetAnnualized(total, supplies, rate)
	fmt.Printf("%+v\n%f", annualized, a)

}

func TestGetDayZeroClock(t *testing.T) {

	//最满意的一次借贷
	var total float64 = 132000
	var supplies = []float64{21844.84, 21991.03, 22197.36, 22288.26, 22448.82, 22642.73}
	var rate = []float64{1309.44 - 900.24, 1096.79 - 764.05, 796.68 - 547.72, 664.31 - 456.72, 430.91 - 296.26, 223.92 - 153.92}

	annualized, a := GetAnnualized(total, supplies, rate)
	fmt.Printf("%+v\n%f", annualized, a)
}
