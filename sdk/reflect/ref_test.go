/*
@Time   : 2022/6/8 17:00
@Author : ckx0709
@Remark :
*/
package reflect

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestGetSliRevByIndex2(t *testing.T) {
	cat := "tom"

	ptr := (*string)(unsafe.Pointer(&cat))
	fmt.Println("ptr:", *ptr)
}

func TestGetSliRevByIndex(t *testing.T) {
	tmp := []string{"a", "b", "c"}
	//tmp := []int{4, 2, 3, 4, 6, 7, 8}
	ptr := GetSliRevByIndex(tmp, 2)

	pointer := *(*string)(ptr)
	fmt.Println("str:", pointer)

	//atomic.LoadPointer(&ptr)

	/*	valType := val.Type()
		fmt.Println(valType)
		ren := reflect.New(valType)
		unsafePointer := ren.UnsafePointer()
		atomic.StorePointer(&unsafePointer, val.UnsafePointer())
		pointer2 := *(*string)(unsafePointer)
		fmt.Printf("ren :%+v\tpointer2 :%s\n", ren, pointer2)*/

}
