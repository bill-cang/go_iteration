/*
@Time   : 2021/12/20 14:14
@Author : ckx0709
@Remark : 循环内部执行defer语句并不是一个好的习惯，此处仅为示例
*/
package unit1


/*
因为是闭包，在for迭代语句中，每个defer语句延迟执行的函数引用的都是同一个i迭代变量，在循环结束后这个变量的值为3，因此最终输出的都是3。
*/
func Clos0() {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}
}

/*
循环体内部再定义一个局部变量，这样每次迭代defer语句的闭包函数捕获的都是不同的变量，这些变量的值对应迭代时的值.
*/
func Clos1() {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func() {
			println(i)
		}()
	}
}

/*
将迭代变量通过闭包函数的参数传入，defer语句会马上对调用参数求值
*/
func Clos2() {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) {
			println(i)
		}(i)
	}
}
