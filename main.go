package main

import (
	"fmt"

	"github.com/jhony/zoo-golang/entrantpkg"
	"github.com/jhony/zoo-golang/speciepkg"
)

type TEntrant = entrantpkg.TEntrant
type TAnimalOptions = speciepkg.TAnimalOptions

func print1(a interface{}) {
	fmt.Printf("%+v\n\n", a)
}

func print2(a, b interface{}) {
	fmt.Printf("%v\n\n%v", a, b)
}

var speciesIds = []string{
	"0938aa23-f153-4937-9f88-4858b24d6bce",
	"e8481c1d-42ea-4610-8e11-1752cfc05a46",
}

var entrants = []TEntrant{
	{Name: "Lara Carvalho", Age: 5},
	{Name: "Frederico Moreira", Age: 5},
	{Name: "Pedro Henrique Carvalho", Age: 5},
	{Name: "Maria Costa", Age: 18},
	{Name: "Núbia Souza", Age: 18},
	{Name: "Carlos Nogueira", Age: 50},
}

// var employe, employes = employeepkg.GetEmployeeByName("nigel")
// var employeBool, employeIds = employeepkg.IsManager("0e7b460e-acf4-4e17-bcb3-ee472265db83")
// var employes, err = employeepkg.GetRelatedEmployees(
//
//	"0e7b460e-acf4-4e17-bcb3-ee472265db83",
//
// )
type TMapS = speciepkg.TMapS

// var s = TMapS{"Specie": "tigers", "Sex": "female"}

// var a = speciepkg.CountAnimals(s)

// var a = entrantpkg.CalculateEntry(entrants)

var s = TAnimalOptions{IncludeNames: true, Sex: "male", Sorted: true}
var a = speciepkg.GetAnimalMap(s)

func main() {
	// print2()
	print1(a)
}
