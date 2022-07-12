/*
@Time   : 2022/6/8 16:51
@Author : ckx0709
@Remark :splic 相关
*/
package reflect

import (
	"reflect"
	"unsafe"
)

//
func GetSliRevByIndex(array interface{}, ind int) (ptr unsafe.Pointer) {
	rev := reflect.ValueOf(array)
	if rev.Kind() != reflect.Slice {
		return reflect.ValueOf(array).UnsafePointer()
	}

	return rev.Slice(ind, ind+1).UnsafePointer()
}
