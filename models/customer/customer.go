package customer

import "time"

type Customer struct {
	Name        string    `json:"name" xml:"name"`
	DateOfBirth time.Time `json:"date_of_birth" xml:"date_of_birth"`
}

func NewCustomer(name string, dateOfBirth time.Time) Customer {
	return Customer{name, dateOfBirth}
}

func (c *Customer) GetName() string {
	return c.Name
}

func (c *Customer) GetAge() string {
	return time.Since(c.DateOfBirth).String()
}
