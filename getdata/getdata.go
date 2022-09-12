package getdata

import (
	"encoding/json"
	"os"
)

type TZoo_data struct {
	Species   []TSpecie  `json:"species"`
	Employees []TEmploye `json:"employees"`
	Hours     THours     `json:"hours"`
	Prices    TPrices    `json:"prices"`
}

type TSpecie struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Popularity   int         `json:"popularity"`
	Location     string      `json:"location"`
	Availability []string    `json:"availability"`
	Residents    []TResident `json:"residents"`
}

type TResident struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

type TEmploye struct {
	Id             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

type THours struct {
	Tuesday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Wednesday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Thursday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Friday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Saturday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Sunday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Monday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
}

type TPrices struct {
	Adult  int `json:"adult"`
	Senior int `json:"senior"`
	Child  int `json:"child"`
}

func GetZoo() TZoo_data {
	data, error := os.ReadFile("./data/zoo_data.json")
	if error != nil {
		panic(error)
	}
	var zoo TZoo_data
	json.Unmarshal([]byte(data), &zoo)
	return zoo
}
