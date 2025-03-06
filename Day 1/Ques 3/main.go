package main

import "fmt"

type Employee interface {
	CalculateSalary() int64
}

type FullTime struct {
	monthlySalary int64
}

type Contractor struct {
	monthlySalary int64
}

type Freelancer struct {
	hoursWorked   int64
	earningsPerHour int64
}

func (f *FullTime) CalculateSalary() int64 {
	return f.monthlySalary
}

func (c *Contractor) CalculateSalary() int64 {
	return c.monthlySalary
}

func (f *Freelancer) CalculateSalary() int64 {
	return f.hoursWorked * f.earningsPerHour
}

func main() {
	fullTimeEmployee := &FullTime{monthlySalary: 15000}
	contractorEmployee := &Contractor{monthlySalary: 3000}
	freelancerEmployee := &Freelancer{hoursWorked: 20, earningsPerHour: 2000}

	employees := []Employee{fullTimeEmployee, contractorEmployee, freelancerEmployee}

	for _, employee := range employees {
		fmt.Printf("Salary: %d\n", employee.CalculateSalary())
	}
}
