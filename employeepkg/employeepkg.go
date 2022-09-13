package employeepkg

import (
	"regexp"
)

func GetEmployeeByName(search string) (TEmploye, []TEmploye) {
	employeFullName := arrOfEmployeeNamesGenerator(Employees)
	employeIdArr := searchEmployee(search, employeFullName)
	var employeesResult []TEmploye
	for _, id := range employeIdArr {
		for _, e := range Employees {
			if id == e.Id {
				employeesResult = append(employeesResult, e)
			}
		}
	}
	return employeesResult[0], employeesResult
}

func searchEmployee(search string, employeFullName []TEmployeFullName) []TEmployeId {
	var employeIdArr []TEmployeId
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
