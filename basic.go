package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包内变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// := 声明变量 不需要var关键字，只能在函数内使用
	a, b, c, s := 3, 4, true, "def"
	d := 5
	fmt.Println(a, b, c, s, d)
}

func eulaer() {
	c := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%.3f\n", c)
}

func triangle() {
	var a, b int = 2, 5
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	// const数值可以作为各种类型使用
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)

	//b,kb,mb,gb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	eulaer()
	triangle()
	consts()
	enums()
}
