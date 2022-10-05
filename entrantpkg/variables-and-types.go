package entrantpkg

import (
	"github.com/jhony/zoo-golang/getdata"
)

var Prices = getdata.GetZoo().Prices

type TEntrant struct {
	Name string
	Age  int
}

type TSortedEntry struct {
	Child  float64
	Adult  float64
	Senior float64
}
