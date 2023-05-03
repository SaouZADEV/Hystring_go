package services

/*
import (
	"database/sql"
	"goredis/errs"
	"goredis/logs"
	"goredis/repositories"
	"log"
	"net/http"
)

type customerService struct {
	customerServiceRepository repositories.CustomerRepository
}

func NewCustomerService(customerServiceRepository repositories.CustomerRepository) customerService {
	return customerService{customerServiceRepository: customerServiceRepository}
}

func (s customerService) GETCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerServiceRepository.GetAll()
	log.Println("Get All customers")
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			Id:           customer.Id,
			Name:         customer.Name,
			Phone_number: customer.Phone_number,
			Date_created: customer.Date_created,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil

}

func (s customerService) GETCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerServiceRepository.GetById(id)
	log.Println("Get customers")
	if err != nil {
		logs.Error(err)
		if err == sql.ErrNoRows {
			return nil, errs.AppError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}
		return nil, errs.AppError{
			Code:    http.StatusBadGateway,
			Message: err.Error(),
		}
	}

	customerResponse := &CustomerResponse{
		Id:           customer.Id,
		Name:         customer.Name,
		Phone_number: customer.Phone_number,
		Date_created: customer.Date_created,
	}

	return customerResponse, nil
}

func (s customerService) ADDCustomer(c repositories.Customer) (int, error) {
	customerid, err := s.customerServiceRepository.AddCustomer(c)
	log.Println("Post add customers")
	if err != nil {
		logs.Error(err)
		return 0, errs.AppError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		}
	}
	return customerid, nil
}

func (s customerService) DELETECustomer(id int) error {
	err := s.customerServiceRepository.DeleteCustomer(id)
	log.Println("Delete customers")
	if err != nil {
		logs.Error(err)
		return errs.AppError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}

func (s customerService) UPDATECustomer(c repositories.Customer, id int) (int, error) {
	customerId, err := s.customerServiceRepository.UpdateCustomer(c, id)
	log.Println("Uplete customers")
	if err != nil {
		logs.Error(err)
		return 0, errs.AppError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	}
	return customerId, nil
}
*/
