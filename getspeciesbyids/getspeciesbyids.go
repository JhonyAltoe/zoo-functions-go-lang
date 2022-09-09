package getspeciesbyids

import (
	"github.com/jhony/zoo-golang/getdata"
)

var Species = getdata.GetZoo().Species

type Tspecie = getdata.Specie

func verifyIfExist(vId string, id []string) bool {
	for _, v := range id {
		if vId == v {
			return true
		}
	}
	return false
}

func GetSpeciesByIds(ids []string) []Tspecie {
	var result []Tspecie

	for _, v := range Species {
		if verifyIfExist(v.Id, ids) {
			result = append(result, v)
		}
	}

	return result
}
