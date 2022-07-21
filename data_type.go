package main

import "fmt"

func main() {
	// boolean types
	var possible bool = true
	fmt.Println(possible)

	// integer types
	var a uint8 = 255
	fmt.Println(a)

	var b uint16 = 65535
	fmt.Println(b)
	
	var c uint32 = 4294967295
	fmt.Println(c)
	
	var d uint64 = 18446744073709551615
	fmt.Println(d)

	var e int8 = -127
	fmt.Println(e)

	var f int16 = -32768
	fmt.Println(f)
	
	var g int32 = -2147483648
	fmt.Println(g)
	
	var h int64 = -9223372036854775808
	fmt.Println(h)

	// floating types
	var i float32
	fmt.Println(i)

	var j float64
	fmt.Println(j)

	var k complex64 = 2+3i
	fmt.Println(k)

	var l complex128 = 20+5i
	fmt.Println(l)

	// other numeric types
	var m byte
	fmt.Println(m)

	var n rune
	fmt.Println(n)

	var o uint
	fmt.Println(o)

	var p int
	fmt.Println(p)

	var q uintptr
	fmt.Println(q)

	// string value
	var text string = "string value"
	fmt.Println(text);
}