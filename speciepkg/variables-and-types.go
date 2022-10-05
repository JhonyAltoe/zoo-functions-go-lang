package speciepkg

import (
	"github.com/jhony/zoo-golang/getdata"
)

var Species = getdata.GetZoo().Species

type TSpecie = getdata.TSpecie
type TResident = getdata.TResident
type TM map[string]int
type MIKE map[string]string
