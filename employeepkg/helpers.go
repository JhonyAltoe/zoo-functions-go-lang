package employeepkg

import "regexp"

func isManager(employeId string) (bool, []TEmployeId) {
	isManager := false
	var managersIds []TEmployeId
	for _, elem := range Employees {
		for _, mId := range elem.Managers {
			if mId == employeId {
				isManager = true
			}
			ok := true
			for _, mId2 := range managersIds {
				if mId2 == mId {
					ok = false
				}
			}
			if ok {
				managersIds = append(managersIds, mId)
			}
		}
	}
	return isManager, managersIds
}

func searchEmployee(search string, employeFullName []TEmployeFullName) []TEmployeId {
	var employeIdArr []string
	for _, v := range employeFullName {
		vBool, _ := regexp.MatchString("(?i)"+search, v.FullName)
		if vBool {
			employeIdArr = append(employeIdArr, v.Id)
		}
	}
	return employeIdArr
}

func arrOfEmployeeNamesGenerator(employees []TEmploye) []TEmployeFullName {
	var fullNameArr []TEmployeFullName
	for _, v := range employees {
		fullName := v.FirstName + " " + v.LastName
		employeFullName := TEmployeFullName{Id: v.Id, FullName: fullName}
		fullNameArr = append(fullNameArr, employeFullName)
	}
	return fullNameArr
}
