package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/errs"
	"server/logs"
	"server/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type customerService struct {
	customerServiceRepository repositories.CustomerRepository
	redisClient               *redis.Client
}

func NewCustomerServiceRedis(customerServiceRepository repositories.CustomerRepository, redisClient *redis.Client) CustomerService {
	return customerService{customerServiceRepository, redisClient}
}

func (s customerService) GETCustomers() ([]CustomerResponse, error) {

	key := "repository:GetAll"
	var customers []CustomerResponse

	customer_tmp, err := s.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(customer_tmp), &customers)
		if err == nil {
			fmt.Println("Redis from service")
			return customers, nil
		}
	}

	customersDB, err := s.customerServiceRepository.GetAll()
	log.Println("Get All customers")
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customerResponses := []CustomerResponse{}
	for _, customer := range customersDB {
		customerResponse := CustomerResponse{
			Id:           customer.Id,
			Name:         customer.Name,
			Phone_number: customer.Phone_number,
			Date_created: customer.Date_created,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	data, err := json.Marshal(customerResponses)
	if err != nil {
		return nil, err
	}
	err = s.redisClient.Set(context.Background(), key, string(data), time.Second*30).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return customerResponses, nil

}

func (s customerService) GETCustomer(id int) (*CustomerResponse, error) {

	key := "repository:GetById"
	var customer CustomerResponse

	customer_tmp, err := s.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(customer_tmp), &customer)
		if err == nil {
			fmt.Println("Redis from service")
			return &customer, nil
		}
	}

	customerDB, err := s.customerServiceRepository.GetById(id)
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
		Id:           customerDB.Id,
		Name:         customerDB.Name,
		Phone_number: customerDB.Phone_number,
		Date_created: customerDB.Date_created,
	}

	data, err := json.Marshal(customerResponse)
	if err != nil {
		return nil, err
	}
	err = s.redisClient.Set(context.Background(), key, string(data), time.Second*30).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
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
