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
	tmp := []string{"a", "b"}
	//tmp := []int{4, 2, 3, 4, 6, 7, 8}

	ptr := GetSliRevByIndex(tmp, 1)
	str := (unsafe.Pointer)(ptr)

	fmt.Println("str:", str)

}
