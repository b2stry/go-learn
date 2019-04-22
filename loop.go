package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
for,if后面的条件没括号
if条件里也可以定义变量
没有while
switch不需要break，也可以直接switch多个条件
 */
func convertToBin(n int) string {
	result := ""
	// 省略初始条件，相当于while
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBin(5), convertToBin(13))
	printFile("abc.txt")
	forever()
}
