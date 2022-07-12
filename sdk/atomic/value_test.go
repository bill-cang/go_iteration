/*
@Time   : 2022/6/16 14:09
@Author : ckx0709
@Remark :
*/
package atomic

import (
	"log"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestGet(t *testing.T) {
	sman := SupperMan
	sman.Key.Store("伟大的领袖毛主席万岁！")
	val := sman.Key.Load()

	str := val.(string)
	log.Println(str)
}

func TestGet2(t *testing.T) {
	var tmp string
	tmp = "111"
	pointer := unsafe.Pointer(&tmp)
	loadPointer := atomic.LoadPointer(&pointer)
	log.Println(pointer, loadPointer)

	s0 := (*string)(pointer)
	tmp2 := "222"
	atomic.StorePointer(&loadPointer, unsafe.Pointer(&tmp2))
	s1 := *(*string)(loadPointer)

	log.Println(*s0, s1)
}
