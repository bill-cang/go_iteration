/*
@Time   : 2021/10/22 9:41
@Author : ckx0709
@Remark :
*/
package util

//二进制转十进制
func Conv2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}
