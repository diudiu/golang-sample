// bufferiosample project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	strReader := strings.NewReader("hello world")

	bufReader := bufio.NewReader(strReader)

	//Peek返回缓存的一个切片，该切片引用缓存中前n字节数据
	data, _ := bufReader.Peek(5)
	fmt.Println(data, bufReader.Buffered())

	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello ")
	fmt.Fprint(w, "world!\n")
	// 写完后记得提交
	w.Flush()

}
