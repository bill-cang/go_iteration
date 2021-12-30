/*
@Time   : 2021/10/22 9:42
@Author : ckx0709
@Remark : 转换
*/
package util

import (
	"fmt"
	"testing"
)

func TestConv2DEC(t *testing.T) {
	dec := Conv2DEC("1011")
	fmt.Print(dec)
}

func TestConv2DEC2(t *testing.T) {
	/*	fmt.Printf("%.2f\n", 1.479+0.005)
		fmt.Printf("%.2f\n", 1.475+0.005)
		fmt.Printf("%.2f\n", 1.470+0.005)*/

	//fmt.Printf("%.2f\n", FloatPositionRetain(18, 2))
	fmt.Println(FloatPositionRetain(1.4799, 2, 0))
	fmt.Println(FloatPositionRetain(1.4755, 2, 0))
	fmt.Println(FloatPositionRetain(1.4701, 2, 0))

	fmt.Println("****")
	fmt.Println(FloatPositionRetain(1.4799, 2, 1))
	fmt.Println(FloatPositionRetain(1.4755, 2, 1))
	fmt.Println(FloatPositionRetain(1.4710, 2, 1))
	fmt.Println(FloatPositionRetain(1.4701, 2, 1))

}

func TestConv2DEC4(t *testing.T) {
	holidays := GetYearHolidays("2021-03-23")
	fmt.Println(holidays)
}

func TestConv2DEC5(t *testing.T) {
	times := GetTimes(34)
	fmt.Println(times)
}

