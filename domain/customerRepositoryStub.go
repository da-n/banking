package domain

import "github.com/da-n/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	return &s.customers[1], nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Daniel", "Bath", "BA2 5RS", "1978-12-10", "1"},
		{"1002", "Sara", "Bath", "BA2 5RS", "1988-03-12", "1"},
	}
	return CustomerRepositoryStub{customers}
}
