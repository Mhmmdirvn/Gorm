package main

type Person struct {
	Id    int	`json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
