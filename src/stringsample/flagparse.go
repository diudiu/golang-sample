package main

import (
	"flag"
	"fmt"
)

// 解析出来的值赋值给新的变量
func style1() {
	methodPtr := flag.String("method", "default", "method of example")
	valuePtr := flag.Int("value", -1, "value of example")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr)
	fmt.Println(flag.Arg(0))
}

// 解析出来的值赋值给已有变量，常用来赋值给结构体
func style2() {
	var method string
	var value int

	flag.StringVar(&method, "method2", "default", "method of example")
	flag.IntVar(&value, "value2", -1, "value of example")

	flag.Parse()

	fmt.Println(method, value)

	// 打印用法
	fmt.Println("\nUsage: ")
	flag.PrintDefaults()
}
