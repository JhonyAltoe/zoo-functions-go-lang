package getspeciesbyids

import (
	"github.com/jhony/zoo-golang/getdata"
)

var species = getdata.GetZoo().Species

func GSBI(id *string) *getdata.Specie {
	var result getdata.Specie
	test := make([]string, 0)

	if len(id) {
		return test
	} else {
		for _, v := range species {
			if v.Id == id {
				result = v
			}
		}
	}

	return result
}
