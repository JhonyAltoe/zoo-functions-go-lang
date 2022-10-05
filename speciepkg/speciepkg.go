package speciepkg

func CountAnimals(v MIKE) interface{} {
	if v["Specie"] == "" && v != nil {
		panic("the key Specie shoudn't be empty")
	}

	m := make(TM)
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
