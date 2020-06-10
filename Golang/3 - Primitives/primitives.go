package main

import (
	"fmt"
)

func main() {
	
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 1 == 2
	fmt.Println(a, b)

	var m bool // Default to False
	fmt.Println(m)

	var num uint8 = 16 // uint16 uint32 uint64
	fmt.Printf("%v, %T\n", num, num)
	
	c := 10
	d := 3
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	fmt.Println(c % d)

	var e int = 10
	var f int8 = 3
	fmt.Println(e + int(f))

	g := 10 // 1010
	h := 3  // 0011

	fmt.Println(g & h) // AND 			0010 = 2
	fmt.Println(g | h) // OR 			1011 = 11
	fmt.Println(g ^ h) // Exclusive OR 	1001 = 9
	fmt.Println(g &^ h)// AND NOT 		0100 = 8

	i := 8 // 2^3
	fmt.Println(i << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(i >> 3)	// 2^3 / 2^3 = 2^0

	j := 3.14 // float32 float64(by default)
	j = 12.7e72
	j = 2.1E14
	fmt.Printf("%v, %T\n", j, j)

	k := 10.2
	l := 3.7
	fmt.Println(k + l)
	fmt.Println(k - l)
	fmt.Println(k * l)
	fmt.Println(k / l)

	var p complex128 = 1 + 2i // complex64
	fmt.Printf("%v, %T\n", p, p)
	q := 2 + 5.2i // complex 128 by default
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)

	fmt.Printf("%v, %T\n", real(p), real(p))
	fmt.Printf("%v, %T\n", imag(p), imag(p))
	
	var r complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", r, r)

	s := "This is a string"
	// s[2] = "u" You cannot do this
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))
	fmt.Println(s + s)
	b1 := []byte(s)
	fmt.Printf("%v, %T\n", b1, b1)

	var r1 rune = 'a'
	fmt.Printf("%v, %T\n", r1, r1)
}
