package main

import "fmt"

// Go语言的基础组成有几个部分：包声明，引入包，函数，变量，语句&&表达式，注释
// 
// 多个标记：关键字，标识符，常量，字符串，符号
// 行分隔符：一行就代表一个语句结束，无其他分隔符
// 注释：单行注释(//)，多行注释 (/* */)
// 标识符：一个或多个字母(A-Z和a-z)，数字(0-9)，下划线组成的序列，但是第一个字符必须是字幕或者下划线而不能是数字
// 		无效标识符：1ab, case(关键字), a+b
// 字符串连接：通过 + 连接
// 关键字：break, default, func, interface, select, case, defer, go, map, struct
// 		  chan, else, goto, package, switch, continue, for, import, return, var
// 		  append, bool, byte, cap, close, complex, complex64, complex128, uint16
// 		  copy, false, float32, float64, imag, int, int8, int16, uint32, int32,
// 		  int64, iota, len, make, new, nil, panic, uint64, print, println, real,
// 		  recover, string, true, uint, uint8, uintptr

func main() {
	fmt.Println("Hello World!")
}