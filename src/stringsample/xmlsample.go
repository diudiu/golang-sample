package main

import (
	"encoding/xml"
	"fmt"
)

// 人物档案结构体，字段需大写字母开头
type person struct {
	Id   int    `xml:"id,attr"`
	Name string `xml:"名字"`
	Age  int
}

func xml_parse() {
	p := person{Id: 1, Name: "stone", Age: 18}

	// 序列化结构体
	var data []byte
	var err error
	if data, err = xml.MarshalIndent(p, "", "  "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data), "\n")

	// 反序列化为结构体
	p2 := new(person)
	// go的经典写法，使用分号";"把变量作用域控制在if范围内
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
}
