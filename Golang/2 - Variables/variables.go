package main

import (
	"fmt"
	"strconv"
)

//var i int = 42 Lower case variables are scoped to the package
var I int = 24 // Upper case are globally visible

var (
	actorName string = "Alan Lyne"
 	companion string = "Venus Chee"
 	doctorNumber int = 3
 	season int = 11
)

var (
	counter int = 0
)

func main() {
	//Anyting delcared in this function is on accessible in this block
	var i int
	i = 42
	// var j float32 = 42
	k := 42.0
	fmt.Println(i)
	fmt.Printf("%v, %T\n", k, k) // v for value, T for type

	i = 23
	fmt.Println(i)

	//var theURL string = "https://google.com"
	
	var j float32
	j = float32(i) // Convert i to a float32 and assigned to j
	fmt.Printf("%v, %T\n", j, j)

	var a string
	a = strconv.Itoa(i) // Conver i to string using string converter package
	fmt.Printf("%v, %T\n", a, a)
}
