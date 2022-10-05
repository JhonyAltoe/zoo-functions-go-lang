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
