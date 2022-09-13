package employeepkg

import (
	"github.com/jhony/zoo-golang/getdata"
)

var Employees = getdata.GetZoo().Employees

type TEmploye = getdata.TEmploye

type TEmployeFullName struct {
	Id       string
	FullName string
}

type TEmployeId = string
