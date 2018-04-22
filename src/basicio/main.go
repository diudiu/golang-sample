// 实现了一个函数ReadFrom分别从字符串、标准输入、文件三个不同的源读取数据
// 由于这三个源都实现了io.Reader接口的Read方法
// 所以可以很方便的从这三个源读取数据

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	//创建一个切片，用于存放读取的数据，相当于buffer
	p := make([]byte, num)
	// 读取数据，存放在p切片中
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func sampleReadFromString() {
	//NewReader返回一个Reader指针
	//Reader是一个结构体类型，绑定了一个Read方法
	//该Read方法实现了io.Reader的Read接口
	data, _ := ReadFrom(strings.NewReader("from string"), 12)

	fmt.Println(data)
}

func sampleReadFromStdin() {
	fmt.Println("please inut from stdin")

	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

func SampleReadFromFile() {
	file, _ := os.Open("main.go")
	defer file.Close()

	data, _ := ReadFrom(file, 9)
	fmt.Println(data)
}
func main() {
	//	sampleReadFromString()
	//	sampleReadFromStdin()
	SampleReadFromFile()
}
