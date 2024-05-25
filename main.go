package main

import (
	"fmt"
)

func main() {
	// // 声明一个指针变量 p
	// var p unsafe.Pointer

	// // 将 p 转换为 uint32 数组的指针，并通过下标访问第 10 个元素
	// // 然后将它的值设为 0xff0000ff
	// *(*uint32)(unsafe.Pointer(uintptr(p) + 10*unsafe.Sizeof(uint32(0)))) = 0xff0000ff

	type A struct {
		a *int
	}

	x := 1

	s := A{a: &x}
	x = 2
	fmt.Println(*s.a)
}
