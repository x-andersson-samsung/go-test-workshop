package golden

import (
	"time"
)

type Company struct {
	Name        string
	Established time.Time
	Departments []Department
	Revenue     float64
	IsPublic    bool
}

type Department struct {
	ID        int
	Name      string
	Budget    float64
	Manager   Employee
	Employees []Employee
	Projects  []Project
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	Skills    []string
	Contact   ContactInfo
	IsActive  bool
}

type ContactInfo struct {
	Email   string
	Phone   string
	Address Address
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

type Project struct {
	ID           int
	Name         string
	Budget       float64
	Technologies []string
	StartDate    time.Time
	EndDate      time.Time
	IsCompleted  bool
	Team         []int // Employee IDs
}
