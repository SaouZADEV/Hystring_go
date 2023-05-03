package repositories

/*
import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type customerRepositoryDB struct {
	db *gorm.DB
}

func NewCustomerRepositoryDB(db *gorm.DB) CustomerRepository {
	db.AutoMigrate(&Customer{})
	mockData(db)
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	var customers []Customer
	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	var customer Customer
	result := r.db.First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("No customer found with id: %d\n", id)
		}
		return nil, result.Error
	}
	return &customer, nil
}

func (r customerRepositoryDB) DeleteCustomer(id int) error {
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

func (r customerRepositoryDB) AddCustomer(c Customer) (int, error) {
	createDate := time.Now()
	c.Date_created = createDate.Format("2006-01-02 15:04:05")

	result := r.db.Create(&c)
	if result.Error != nil {
		return 0, result.Error
	}
	log.Printf("DEBUG: inserted record with id: %v\n", c.Id)
	return int(c.Id), nil
}

func (r customerRepositoryDB) UpdateCustomer(c Customer, id int) (int, error) {
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
*/
