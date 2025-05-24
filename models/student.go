package models

type Student struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students = []Student{}
