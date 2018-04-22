package main

import (
	"fmt"
	"strings"
)

func str_op() {
	s := "hello world"
	fmt.Println("原始字符串：", s)

	// 是否包含
	fmt.Println("是否包含hello: ", strings.Contains(s, "hello"))
	fmt.Println("是否包含?: ", strings.Contains(s, "?"))

	// 索引， base0
	fmt.Println("o的索引是：", strings.Index(s, "o"))

	ss := "1#2#345"
	fmt.Println("原始字符串：", ss)

	// 切割字符串
	splitedStr := strings.Split(ss, "#")
	fmt.Println("切割后的字符串：", splitedStr)

	// 合并字符串
	fmt.Println("合并后的字符串：", strings.Join(splitedStr, "/"))

	// 前后缀
	fmt.Println("是否有he前缀, ", strings.HasPrefix(s, "he"))
	fmt.Println("是否有ld后缀：", strings.HasSuffix(s, "ld"))
}
