package main

import (
	"fmt"

	"github.com/jhony/zoo-golang/employeepkg"
)

var speciesIds = []string{
	"0938aa23-f153-4937-9f88-4858b24d6bce",
	"e8481c1d-42ea-4610-8e11-1752cfc05a46",
}

var e, es = employeepkg.GetEmployeeByName("a")

func main() {
	// fmt.Printf(
	// 	"%+v\n",
	// 	// speciepkg.GetSpeciesByIds(speciesIds),
	// 	// speciepkg.GetAnimalsOlderThan("lions", 8),
	// )
	fmt.Printf(
		"%v\n\n%v",
		e, es,
	)
}
