/*
@Time   : 2021/12/30 9:35
@Author : ckx0709
@Remark : 年假计算器
*/
package util

import "time"

/*
获取应得年假天数，年假基数5
上一年度入职日历天数/365天*全年应享受天数
*/
func GetYearHolidays(inductionDay string) float64 {
	parse, err := time.Parse("2006-01-02", inductionDay)
	if err != nil {
		return 0.0
	}

	day := parse.YearDay()
	return float64((365-day)*5+1) / 365
}
