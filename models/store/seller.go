package store

import "time"

type Seller struct {
	Name        string    `json:"name" xml:"name"`
	DateOfBirth time.Time `json:"date_of_birth" xml:"date_of_birth"`
}

func NewSeller(name string, dateOfBirth time.Time) Store {
	return &Seller{name, dateOfBirth}
}

func (c *Seller) GetName() string {
	return c.Name
}

func (c *Seller) GetAge() string {
	return time.Since(c.DateOfBirth).String()
}
