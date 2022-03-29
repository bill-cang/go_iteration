/*
@Time   : 2022/3/29 18:22
@Author : ckx0709
@Remark :
*/
package constants

import (
	"fmt"
	"testing"
)

func TestGetRandUserAgent(t *testing.T) {
	for i := 0; i < 20; i++ {
		agent := GetRandUserAgent()
		fmt.Println(agent)
	}
}
