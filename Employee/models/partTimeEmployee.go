package models

import "fmt"

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked uint64
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d, HoursWorked: %d", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}
