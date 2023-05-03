package handlers

/*
import (
	"encoding/json"
	"fmt"
	"goredis/errs"
	"goredis/repositories"
	"goredis/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) customerHandler {
	return customerHandler{customerService: customerService}
}

func (h customerHandler) GetCustomers(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)
	customers, err := h.customerService.GETCustomers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	json_, err := json.Marshal(customers)
	_ = json_
	if err != nil {
		return err
	}
	return c.JSON(customers)
}

func (h customerHandler) GetCustomer(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)
	customerID, _ := strconv.Atoi(c.Params("id"))

	customer, err := h.customerService.GETCustomer(customerID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(customer)
}

func (h customerHandler) AddCustomer(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)
	var customer repositories.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	customerID, err := h.customerService.ADDCustomer(customer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Add successfully with ID %d", customerID)})
}

func (h customerHandler) DeleteCustomer(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)
	customerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = h.customerService.DELETECustomer(customerID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Delete successfully with ID %d", customerID)})
}

func (h customerHandler) UpdateCustomer(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)
	customerID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var customer repositories.Customer
	err = c.BodyParser(&customer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	customerID_, err := h.customerService.UPDATECustomer(customer, customerID)
	_ = customerID_
	if err != nil {
		appErr, ok := err.(errs.AppError)
		if ok {
			return c.Status(appErr.Code).SendString(appErr.Message)
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Update successfully with ID %d", customerID)})
}
*/
