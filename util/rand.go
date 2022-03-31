/*
@Time   : 2022/3/29 18:03
@Author : ckx0709
@Remark :
*/
package util

import (
	"math/rand"
	"time"
)

func GetRand(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
