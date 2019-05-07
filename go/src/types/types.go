package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Golang has the following types:

	// Boolean
	a, b := true, false
	fmt.Println("A is", a, "and B is", b)
	c := a && b
	fmt.Println("C is", c, "as the result of a && b")
	d := a || b
	fmt.Println("D is", d, "as the result of a || b")

	// Signed integers
	// int is the default type for integers. When in doubt, this should be used.
	i := int(100)
	fmt.Printf("i is a variable of type %T and size of %d\n", i, unsafe.Sizeof(i))
	// i8 is int8, a representation of 8 bit signed integer. int8 goes from -128
	// to 127.
	i8 := int8(126)
	fmt.Printf("i8 is a variable of type %T and size of %d\n", i8, unsafe.Sizeof(i8))
	// i16 is int16, a representation of 16 bit signed integer. int16 goes from 
	// -32768 to 32767.
	i16 := int16(32766)
	fmt.Printf("i16 is a variable of type %T and size of %d\n", i16, unsafe.Sizeof(i16))
	// i32 is int32, a representation of 32 bit signed integer. int32 goes from
	// -2147483648 to 2147483647.
	i32 := int32(2147483646)
	fmt.Printf("i32 is a variable of type %T and size of %d\n", i32, unsafe.Sizeof(i32))
	// i64 is int64, a representation of 64 bit signed integer. int64 goes from
	// -9223372036854775808 to 9223372036854775807.
	i64 := int64(9223372036854775806)
	fmt.Printf("i64 is a variable of type %T and size of %d\n", i64, unsafe.Sizeof(i64))

	// Unsigned integers
	// ui8 is int8, a representation of 8 bit unsigned integer. uint8 goes from 0
	// to 255.
	ui8 := uint8(254)
	fmt.Printf("ui8 is a variable of type %T and size of %d\n", ui8, unsafe.Sizeof(ui8))
	// i16 is int16, a representation of 16 bit signed integer. uint16 goes from 0
	// to 65535.
	ui16 := uint16(65534)
	fmt.Printf("ui16 is a variable of type %T and size of %d\n", ui16, unsafe.Sizeof(ui16))
	// i32 is int32, a representation of 32 bit signed integer. uint32 goes from 0
	// to 4294967295.
	ui32 := uint32(4294967294)
	fmt.Printf("ui32 is a variable of type %T and size of %d\n", ui32, unsafe.Sizeof(ui32))
	// i64 is int64, a representation of 64 bit signed integer. uint64 goes from 0
	// to 18446744073709551615.
	ui64 := uint64(18446744073709551614)
	fmt.Printf("ui64 is a variable of type %T and size of %d\n", ui64, unsafe.Sizeof(ui64))

	// Floating points
	// Notice that there is no float as default type.
	// float32 uses 32 bits.
	f32 := float32(5.67)
	fmt.Printf("f32 is a variable of type %T and size of %d\n", f32, unsafe.Sizeof(f32))
	// float64 is the default type for all floating point values. It uses 64 bits,
	// so it takes twice the size and math operations may consume more memory than
	// usual.
	f64 := float64(5.67)
	fmt.Printf("f64 is a variable of type %T and size of %d\n", f64, unsafe.Sizeof(f64))

	// Complex types
	comp64 := complex(3, 4)
	comp
}
