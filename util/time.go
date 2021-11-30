/*
@Time   : 2021/11/30 15:16
@Author : ckx0709
@Remark : 时间工具
*/
package util

import "time"

//获取anyDay天凌晨
func GetDayZeroClock(anyDay int64) time.Time {
	n := time.Now().Add(time.Hour * time.Duration(anyDay*24))
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
}
