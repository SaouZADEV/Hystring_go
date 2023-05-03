package services

import "server/repositories"

type CustomerResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Phone_number string `json:"phone_number"`
	Date_created string `json:"date_created"`
}

type CustomerService interface {
	GETCustomers() ([]CustomerResponse, error)
	GETCustomer(int) (*CustomerResponse, error)
	UPDATECustomer(repositories.Customer, int) (int, error)
	ADDCustomer(repositories.Customer) (int, error)
	DELETECustomer(int) error
}
