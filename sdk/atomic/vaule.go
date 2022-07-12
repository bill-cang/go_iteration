/*
@Time   : 2022/6/16 14:07
@Author : ckx0709
@Remark :
*/
package atomic

import "sync/atomic"

var SupperMan struct {
	Key atomic.Value
}
