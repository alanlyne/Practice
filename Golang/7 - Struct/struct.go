package main

import(
	"fmt"
	"reflect"
)
type Doctor struct {
	number int
	actorName string
	companions[] string
}


type Animal struct {
	Name string `required max:"100"` // Back ticks
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

func main() {

	aDoctor := Doctor{
		number: 3,
		actorName: "Alan Lyne",
		companions: []string {
			"Jack Lyne",
			"Venus Chee",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	aCyclist := struct{name string}{name: "Chris Froome"} // Anonymous Struct
	fmt.Println(aCyclist)
	anotherCyclist := aCyclist // Independant copy of data set
	anotherCyclist.name = "Peter Sagan"
	fmt.Println(aCyclist)
	fmt.Println(anotherCyclist)

	yetAnotherCyclist := &aCyclist // Point to same underlying data
	yetAnotherCyclist.name = "Dan Martin"
	fmt.Println(aCyclist)
	fmt.Println(yetAnotherCyclist)


	// b := Bird{} Embedding
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false

	b := Bird {  // Also Embedding
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly: false,
	}
	fmt.Println(b)
	fmt.Println(b.Name)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}