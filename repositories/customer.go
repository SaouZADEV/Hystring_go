package repositories

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Phone_number string `db:"phone_number" json:"phone_number"`
	Date_created string `db:"date_created" json:"date_created"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
	UpdateCustomer(Customer, int) (int, error)
	AddCustomer(Customer) (int, error)
	DeleteCustomer(int) error
}

func mockData(db *gorm.DB) error {

	var count int64
	db.Model(&Customer{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	customers := []Customer{}
	for i := 0; i < 500; i++ {
		createDate := time.Now()

		customers = append(customers, Customer{
			Name:         fmt.Sprintf("Customer%v", i+1),
			Phone_number: strconv.Itoa(random.Intn(100)),
			Date_created: createDate.Format("2006-01-02 15:04:05"),
		})
	}

	return db.Create(&customers).Error
}
