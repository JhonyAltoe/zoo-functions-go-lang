package speciepkg

import (
	"github.com/jhony/zoo-golang/getdata"
)

var Species = getdata.GetZoo().Species

type TSpecie = getdata.TSpecie
type TResident = getdata.TResident
type TMapI map[string]int
type TMapS map[string]string

// type TSpecieToMap map[string]string
type TAnimalMapName map[string]map[string][]string

var Test = TAnimalMapName{
	"NE": nil, "NW": nil, "SE": nil, "SW": nil,
}

type TAnimalMap = map[string][]string

type TAnimalOptions = struct {
	IncludeNames bool
	Sex          string
	Sorted       bool
}
