package main

import . "github.com/NursultanNurgaliyev/Assignment1/Employee/models"

func main() {
	c := NewCompany()
	c.AddEmployee(FullTimeEmployee{ID: 1, Name: "Alice", Salary: 5000})
	c.AddEmployee(PartTimeEmployee{ID: 2, Name: "Bob", HourlyRate: 20, HoursWorked: 100})

	c.ListEmployees()
}
