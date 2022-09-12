package speciepkg

func stringVerifyIntoArr(baseArr []string, s string) bool {
	for _, v := range baseArr {
		if s == v {
			return true
		}
	}
	return false
}

func verifyIfResidentHasMinAge(residents []TResident, age int) bool {
	var boolVar = true
	for _, v := range residents {
		if v.Age < age {
			boolVar = false
		}
	}
	return boolVar
}

func GetSpeciesByIds(ids []string) []Tspecie {
	var result []Tspecie

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
