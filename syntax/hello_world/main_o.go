package main

import (
	_ "fmt" // 匿名引用
	//	"unicode/utf8"
)

func main() {
	/*
		println("hello world")
		println(len("hello world"))           // 输出11
		println(len("你好"))                    // 输出6
		println(utf8.RuneCountInString("你好")) // 输出2
		println("你好")                         // 输出2
	*/
	for i := 0; i < len("hello world"); i++ {
		println("hello world i", i) // 输出11
	}
	a, b, c := 1, 2, "hello"
	println(a, b, c)
}
