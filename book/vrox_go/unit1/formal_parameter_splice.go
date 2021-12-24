/*
@Time   : 2021/12/20 14:34
@Author : ckx0709
@Remark :
*/
package unit1

/*func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}*/

type IntSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

func twice(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}
