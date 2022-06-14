/*
@Time   : 2022/6/8 16:51
@Author : ckx0709
@Remark :splic 相关
*/
package reflect

import (
	"reflect"
)

//
func GetSliRevByIndex(array interface{}, ind int) (ptr uintptr) {
	rev := reflect.ValueOf(array)
	if rev.Kind() != reflect.Slice {
		return reflect.Zero(nil).UnsafeAddr()
	}

	return (*reflect.SliceHeader)(rev.Slice(ind, ind+1).UnsafePointer()).Data
}
