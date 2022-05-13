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

var TestVariable = "abcdefg"

type IntSliceHeader struct {
	Data []int `gorm:"column:data"`
	Len  int   `gorm:"column:len"`
	Cap  int   `gorm:"column:cap"`
}

func twice(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}
