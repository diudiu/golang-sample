// readimageinfo project main.go
package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("penguin.bmp")
	defer file.Close()

	// 字节序 linux、windows通常是小端的
	var headA, headB byte
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	var size uint
	binary.Read(file, binary.LittleEndian, &size)

	var reservedA, reservedB byte
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	var offbits uint
	binary.Read(file, binary.LittleEndian, &offbits)

	//	fmt.Printf("%c%c /n", headA, headB)

}
