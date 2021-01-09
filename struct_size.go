package main

import (
	"fmt"
	"unsafe"
)

// Alignment for user defined types has a huge role
// in memory allocation. As per CPU size(32/64) also called as
// word size and type of variable inside struct, size of struct varies.

// ex1: Memory allocation map:
// a = 1 byte
// padding = 7 bytes (since next int64 will start from 8)
// c = 8 bytes
// b = 4 bytes
// padding of 4 bytes to make it compatible as per word size
// i.e. total size should be divisible by word size
// memory plot : bool |0-1| padding |1-8| int64 |8-16| float32 |16-20| padding |20-24|
type struct1 struct {
	a bool
	c int64
	b float32
}

// ex2: Memory allocation map:
// c = 8 bytes
// b = 4 bytes
// c = 1 byte
// padding = 3 bytes (8+4+1 not divisible by word size(here 8) so add +3)
// memory plot : int64 |0-8| float32 |8-12| bool |12-13| padding |13-16|
type struct2 struct {
	c int64
	b float32
	a bool
}

func main() {
	var s1 struct1
	var s2 struct2
	fmt.Println("Size of s1: ", unsafe.Sizeof(s1))
	fmt.Println("Size of s2: ", unsafe.Sizeof(s2))
}
