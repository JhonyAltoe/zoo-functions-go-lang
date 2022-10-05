package employeepkg

import (
	"errors"
)

func GetRelatedEmployees(employeId string) ([]string, error) {
	isManager, _ := isManager(employeId)
	var employeesArr []string
	if isManager {
		for _, e := range Employees {
			for _, m := range e.Managers {
				if m == employeId {
					employeesArr = append(employeesArr, e.FirstName+" "+e.LastName+",")
				}
			}
		}
		return employeesArr, nil
	}

	return employeesArr, errors.New("o id inserido não é de um colaborador gerente")
}

func GetEmployeeByName(search string) (TEmploye, []TEmploye) {
	employeFullName := arrOfEmployeeNamesGenerator(Employees)
	employeIdArr := searchEmployee(search, employeFullName)
	var employeesResult []TEmploye
	for _, e := range Employees {
		for _, id := range employeIdArr {
			if id == e.Id {
				employeesResult = append(employeesResult, e)
			}
		}
	}
	return employeesResult[0], employeesResult
}
