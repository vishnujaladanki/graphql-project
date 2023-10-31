// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `json:"Name"`
	Location string `json:"Location"`
}

type Job struct {
	gorm.Model
	JobID       int    `gorm:"primaryKey;autoIncrement"`
	Tittle      string `json:"Tittle"`
	Description string `json:"Description"`
	CompanyID   int    `json:"CompanyId"`
}

type NewCompany struct {
	Name     string `json:"Name"`
	Location string `json:"Location"`
}

type NewJob struct {
	Tittle      string `json:"Tittle"`
	Description string `json:"Description"`
	CompanyID   int    `json:"CompanyId"`
}

type NewUser struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type User struct {
	gorm.Model
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}
