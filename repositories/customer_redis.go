package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type customerRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewCustomerRepositoryRedis(db *gorm.DB, redisClient *redis.Client) CustomerRepository {
	db.AutoMigrate(&Customer{})
	mockData(db)
	return customerRepositoryRedis{db, redisClient}
}

func (r customerRepositoryRedis) GetAll() ([]Customer, error) {

	key := "repository:GetAll"
	var customers []Customer

	customer_tmp, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(customer_tmp), &customers)
		if err == nil {
			fmt.Println("Redis from repository")
			return customers, nil
		}
	}

	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	data, err := json.Marshal(customers)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*30).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return customers, nil
}

func (r customerRepositoryRedis) GetById(id int) (*Customer, error) {

	var customer Customer
	key := "repository:GetById"

	customer_tmp, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(customer_tmp), &customer)
		if err == nil {
			fmt.Println("Redis")
			return &customer, nil
		}
	}

	result := r.db.First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("No customer found with id: %d\n", id)
		}
		return nil, result.Error
	}

	data, err := json.Marshal(&customer)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*30).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return &customer, nil
}

func (r customerRepositoryRedis) DeleteCustomer(id int) error {
	result := r.db.Delete(&Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("customer with id %d not found", id)
	}
	log.Printf("DEBUG: successfully deleted record with id: %d\n", id)
	return nil
}

func (r customerRepositoryRedis) AddCustomer(c Customer) (int, error) {
	createDate := time.Now()
	c.Date_created = createDate.Format("2006-01-02 15:04:05")

	result := r.db.Create(&c)
	if result.Error != nil {
		return 0, result.Error
	}
	log.Printf("DEBUG: inserted record with id: %v\n", c.Id)
	return int(c.Id), nil
}

func (r customerRepositoryRedis) UpdateCustomer(c Customer, id int) (int, error) {
	result := r.db.Model(&Customer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":         c.Name,
		"phone_number": c.Phone_number,
	})
	if result.Error != nil {
		if result.RowsAffected == 0 {
			log.Println("DEBUG: customer id not found in the table")
			return 0, fmt.Errorf("customer with id %d not found", id)
		}
		return 0, result.Error
	}
	log.Printf("DEBUG: customer with id %d updated successfully\n", id)
	return int(result.RowsAffected), nil
}
