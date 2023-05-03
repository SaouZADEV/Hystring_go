package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/errs"
	"server/repositories"
	"server/services"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService services.CustomerService
	redisClient     *redis.Client
}

func NewCustomerHandlerRedis(customerService services.CustomerService, redisClient *redis.Client) CustomerHandler {
	return CustomerHandler{customerService, redisClient}
}

func (h CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)

	key := "repository:GetAll"

	responseJson, err := h.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		fmt.Println("redis from handler")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	customers, err := h.customerService.GETCustomers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	json_, err := json.Marshal(customers)
	_ = json_
	if err != nil {
		return err
	}

	h.redisClient.Set(context.Background(), key, string(json_), time.Second*10)

	fmt.Println("database")

	return c.JSON(customers)
}

func (h CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	ipAddress := c.IP()
	fmt.Printf("API accessed from IP: %s\n", ipAddress)
	customerID, _ := strconv.Atoi(c.Params("id"))

	customer, err := h.customerService.GETCustomer(customerID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(customer)
}

func (h CustomerHandler) AddCustomer(c *fiber.Ctx) error {
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

func (h CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
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

func (h CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
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
