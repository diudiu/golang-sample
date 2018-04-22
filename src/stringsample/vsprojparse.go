// https://www.imooc.com/video/10248
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// xml是比较复杂的文本，解析需要用到编译原理知识 -- 把文本流解析成token -- xml.NewDecoder已经封装了此操作
func vsproj_parser() {
	content, err := ioutil.ReadFile("vsproj.csproj")
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	// Token()返回每个xml节点
	var t xml.Token
	var inItemGroup bool // 状态机标识
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local

			if inItemGroup {
				if name == "Compile" {
					fmt.Println(getAttributeValue(token.Attr, "Include"))
				}
			} else {
				if name == "ItemGroup" {
					inItemGroup = true
				}
			}
		case xml.EndElement:
			{
				if inItemGroup {
					if token.Name.Local == "ItemGroup" {
						inItemGroup = false
					}
				}
			}
		}
	}
}

func getAttributeValue(attr []xml.Attr, name string) string {
	for _, a := range attr {
		if a.Name.Local == name {
			return a.Value
		}
	}

	return ""
}
