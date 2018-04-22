package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var line int
	reader := bufio.NewReader(file)
	for {
		// 当读取的行超过buffer大小，isPrefix设为true，直到该行全部读取完，才设置成false
		_, isPrefix, err := reader.ReadLine()

		if err != nil {
			break
		}

		if !isPrefix {
			line++
		}
	}
	fmt.Println(line)
}
