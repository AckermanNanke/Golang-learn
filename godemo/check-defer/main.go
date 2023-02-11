package checkDefer

import (
	"fmt"
)

/*
*	闭包修改外部变量，x已修改但不会影响返回值
* X已复制给返回变量
 */
func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Println("f1", x)
	}()
	return x
}

/*
*	闭包修改外部变量
*	return时： 5 赋值给变量X
*	defer时： X 自增一
*	返回6
 */
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

/*
*	闭包修改外部变量
*	return时： X 赋值给变量Y
*	defer时： X 自增一
*	返回变量Y 即5
 */
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

/*
*	闭包修改外部变量
*	return时： X 赋值给变量Y
*	defer时： 函数参数使用变量X（defer 函数时即确定了其参数的值因此X由 0 自增一），则修改的是自己作用域内的X，而不是外部的
*	返回5
 */
func f4() (x int) {
	defer func(x int) {
		x++
		fmt.Println("f4", x)
	}(x)
	return 5
}

func RunDefer() {
	fmt.Println("\n", "checkDefer=============================")
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println("checkDefer=============================", "\n")
}
