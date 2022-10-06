package speciepkg

import "sort"

func GetAnimalMap(options TAnimalOptions) interface{} {
	var aMapName = make(TAnimalMapName)
	var aMap = make(TAnimalMap)
	for _, specie := range Species {
		if options.IncludeNames {
			aMapName[specie.Location] = map[string][]string{}
		} else {
			aMap[specie.Location] = []string{}
		}
		for _, resident := range specie.Residents {
			if resident.Sex == options.Sex || len(options.Sex) == 0 {
				var toSort []string
				if locN := aMapName[specie.Location]; options.IncludeNames {
					locN[specie.Name] = append(locN[specie.Name], resident.Name)
					toSort = locN[specie.Name]
				} else {
					aMap[specie.Location] = append(aMap[specie.Location], resident.Name)
					toSort = aMap[specie.Location]
				}
				if options.Sorted {
					sort.Strings(toSort)
				}
			}
		}
	}
	if options.IncludeNames {
		return aMapName
	}
	return aMap
}

func CountAnimals(v TMapS) interface{} {
	if v["Specie"] == "" && v != nil {
		panic("the key Specie shoudn't be empty")
	}

	m := make(TMapI)
	for _, e := range Species {
		if v == nil {
			m[e.Name] = len(e.Residents)
			continue
		}

		if e.Name == v["Specie"] {
			if v["Sex"] == "male" || v["Sex"] == "female" {
				var count int
				for _, r := range e.Residents {
					if r.Sex == v["Sex"] {
						count++
					}
				}
				m[e.Name] = count
				break
			}
			m[e.Name] = len(e.Residents)
			break
		}
	}
	return m
}

func GetSpeciesByIds(ids []string) []TSpecie {
	var result []TSpecie

	for _, v := range Species {
		if stringVerifyIntoArr(ids, v.Id) {
			result = append(result, v)
		}
	}

	return result
}

func GetAnimalsOlderThan(name string, minAge int) bool {
	var boolVar bool
	for _, v := range Species {
		if v.Name == name {
			boolVar = verifyIfResidentHasMinAge(v.Residents, minAge)
		}
	}
	return boolVar
}
