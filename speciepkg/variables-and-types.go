package speciepkg

import (
	"github.com/jhony/zoo-golang/getdata"
)

var Species = getdata.GetZoo().Species

type Tspecie = getdata.TSpecie
type TResident = getdata.TResident
