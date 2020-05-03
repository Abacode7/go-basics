package main

import "fmt"

var var1 int
var var2 int8
var var3 int16 = 30
var var4 int32 = 45
var var5 int64 = 10000000
var var7, var8, var9 = 1, false, 3.3

// ALPHA variable is to be exported
const ALPHA = 789

var(
	var10 = 4.5
	a1, b1 float32 = 4.5, 7.0
	c1 float64
	d1 complex64
	e1 complex128
)

const(
	a = 1 // a == 1
	b = 2 // b == 2
	c     // c == 2
	d     // d == 2
	e = iota // e == 2
    f        // f == 3 (implicitely d = iota)
)


func main()  {
	// type declaration allowed in functions
	var6 := 122.0

	//rune is an alias for int32
	var var11 rune = 578 

	//byte is an alias for int8
	var var12 int8 = 75

	fmt.Printf("the value is %v and the type is %T\n", var1, var1)
	fmt.Printf("the value is %v and the type is %T\n", var2, var2)
	fmt.Printf("the value is %v and the type is %T\n", var3, var3)
	fmt.Printf("the value is %v and the type is %T\n", var4, var4)
	fmt.Printf("the value is %v and the type is %T\n", var5, var5)
	fmt.Printf("the value is %v and the type is %T\n", var6, var6)
	fmt.Printf("the value is %v and the type is %T\n", var7, var7)
	fmt.Printf("the value is %v and the type is %T\n", var8, var8)
	fmt.Printf("the value is %v and the type is %T\n", var9, var9)
	fmt.Printf("the value is %v and the type is %T\n", var10, var10)
	fmt.Printf("the value is %v and the type is %T\n", var11, var11)
	fmt.Printf("the value is %v and the type is %T\n", var12, var12)
}