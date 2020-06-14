package main

import(
	"fmt"
)

func main() {

	countyPop := make(map[string]int)
	countyPop = map[string]int{ // Key -> String, Value -> Int
		"Meath": 789281,
		"Kildare": 674927,
		"Louth": 547289,
		"Wicklow": 298408,
		"Cavan": 104876,
	}
	countyPop["Dublin"] = 1809476
	fmt.Println(countyPop)

	fmt.Println(countyPop["Meath"])

	delete(countyPop, "Cavan")
	fmt.Println(countyPop["Cavan"])
	fmt.Println(countyPop)

	pop, ok := countyPop["Cavan"] // Replace "pop" with "_" if we only care about true/false
	fmt.Println(pop, ok)
	pop, ok = countyPop["Kildare"]
	fmt.Println(pop, ok)

	fmt.Println(len(countyPop))

	cp := countyPop // Passed by referenced - milipulates both if changed
	delete(cp, "Louth")
	fmt.Println(cp)
	fmt.Println(countyPop)

}