package getdata

import (
	"encoding/json"
	"os"
)

type Zoo_data struct {
	Species   []Specie  `json:"species"`
	Employees []Employe `json:"employees"`
	Hours     Hours     `json:"hours"`
	Prices    Prices    `json:"prices"`
}

type Specie struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Popularity   int        `json:"popularity"`
	Location     string     `json:"location"`
	Availability []string   `json:"availability"`
	Residents    []Resident `json:"residents"`
}

type Resident struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

type Employe struct {
	Id             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

type Hours struct {
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

type Prices struct {
	Adult  int `json:"adult"`
	Senior int `json:"senior"`
	Child  int `json:"child"`
}

func GetZoo() Zoo_data {
	data, error := os.ReadFile("./data/zoo_data.json")
	if error != nil {
		panic(error)
	}
	var zoo Zoo_data
	json.Unmarshal([]byte(data), &zoo)
	return zoo
}
