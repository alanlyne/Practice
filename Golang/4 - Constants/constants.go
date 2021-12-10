package main

import (
	"fmt"
)

const e float32 = 27.50

const (
	f = iota // Scopped to constant block
	f1 = iota
	f2 = iota
)

const (
	g = iota
	g1
	g2
)

const (
	_ = iota // Write only value (we don't care about it)
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_ = iota
	KB = 1 << (10 * iota) // 1 * 2^10, then 2^100,then 2^1000 and so on
	MB
	GB
	TB
	PB
)

const (
	isAdmin = 1 << iota //00000001
	isHeadquaters 		//00000010
	canSeeFinancials 	//00000100
	canSeeAfrica		//00001000
	canSeeAsia			//00010000
	canSeeEurope		//00100000
	canSeeNAmerica		//01000000
	canSeeSAmerica		//10000000
)

func main() {
	
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const e = 14 //Const e inside of main "wins"
	fmt.Printf("%v, %T\n", e, e)

	var e1 int = 27
	fmt.Printf("%v, %T\n", e + e1, e + e1) // Must be done with same types (When const with diff type is involved)

	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", f1, f1)
	fmt.Printf("%v, %T\n", f2, f2)

	fmt.Println(g, g1, g2)

	var specialistType int = catSpecialist
	fmt.Println(specialistType == catSpecialist)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB) // Two decimal place and "GB" after

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope // Combine all data togther as they are all true
	fmt.Printf("%b\n", roles) // Byte representation of 37 (00100101)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquaters & roles == isHeadquaters)

}
