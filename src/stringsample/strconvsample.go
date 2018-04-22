package main

import (
	"fmt"
	"strconv"
)

func str_conv() {
	// 基本数值转换
	// itoa 输入整型 输出字符串
	fmt.Println(strconv.Itoa(10))
	// atoi ascii -> int
	fmt.Println(strconv.Atoi("711"))

	// 解析
	fmt.Println(strconv.ParseBool("f"))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	// 格式化
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 2))
}
