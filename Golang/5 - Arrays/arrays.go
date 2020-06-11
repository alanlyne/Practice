package main

import (
	"fmt"
)

func main() {

	//ARRAYS

	//grades := [3]int{97, 85, 93} An array on 3 ints
	grades := [...]int{97, 85, 93, 78} // [...] Create an array to hold the data I pass in i.e, 4
	fmt.Println("Grades:", grades)

	var students [3]string
	fmt.Println(students)
	students[0] = "Alan"
	fmt.Println(students)
	students[1] = "Venus"
	students[2] = "Darragh"
	fmt.Println(students)

	fmt.Println("Number of students:", len(students))

	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a // 'b' is assinged to a copy of the array 'a'
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	a1 := [...]int{1, 2, 3}
	b1 := &a1 // 'b1' is pointing to the array 'a1'.
	b1[1] = 5 // When we change the value, we change the underlying data in both
	fmt.Println(a1)
	fmt.Println(b1)

	// SLICES

	m := []int{7, 8, 9}
	fmt.Println(m, "\nLength:", len(m), "\nCapacity:", cap(m))
	n := m // Slices point to the same underlying data
	n[1] = 5
	fmt.Println(m, "\nLength M:", len(m), "\nCapacity M:", cap(m))
	fmt.Println(n, "\nLength N:", len(n), "\nCapacity N:", cap(n))

	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w := v[:] 	// Slice of all elements
	x := v[3:] 	// Slice from 4th element to the end
	y := v[:6]	//  Slice first 6 elements
	z := v[3:6] // Slice the 4th, 5th and 6th element
	v[5] = 42
	fmt.Println(v, "\n", w, "\n", x, "\n", y, "\n", z) // All the above can be done with an array also

	c := make([]int, 3, 100)
	fmt.Println(c)
	fmt.Println("Length C:", len(c))
	fmt.Println("Capacity C:", cap(c))

	d := []int{}
	fmt.Println(d)
	fmt.Println("Length D:", len(d))
	fmt.Println("Capacity D:", cap(d))
	d = append(d, 1, 2, 3, 4, 5)
	d = append(d, []int{6, 7, 8, 9}...) // Concatenate slices
	fmt.Println(d)
	fmt.Println("Length D:", len(d))
	fmt.Println("Capacity D:", cap(d))

	e := d[1:] // Remove first element
	fmt.Println(e)
	e = d[:len(d)-1] // Remove last element
	fmt.Println(e)
	e = append(d[:2], d[3:]...) // Remove 3rd element 
	fmt.Println(e)
	fmt.Println(d) // Original array is messed up
}
