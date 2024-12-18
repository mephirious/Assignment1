package models

import (
	"fmt"
	"strconv"
)

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{make(map[string]Employee)}
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.Employees["FT"+strconv.FormatUint(e.ID, 10)] = emp
	case PartTimeEmployee:
		c.Employees["PT"+strconv.FormatUint(e.ID, 10)] = emp
	}
}

func (c *Company) ListEmployees() {
	for _, employee := range c.Employees {
		fmt.Println(employee.GetDetails())
	}
}
