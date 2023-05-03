package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"server/handlers"
	"server/repositories"
	"server/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	initConfig()
	initTimeZone()
	db := initDatabase()
	redisClient := initRedis()
	customerRepositoryDB := repositories.NewCustomerRepositoryRedis(db, redisClient)
	customerService := services.NewCustomerServiceRedis(customerRepositoryDB, redisClient)
	customerHandler := handlers.NewCustomerHandlerRedis(customerService, redisClient)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
		AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token",
		AllowCredentials: true,
	}))

	app.Get("/customers", customerHandler.GetCustomers)
	app.Get("/customers/:id", customerHandler.GetCustomer)
	app.Post("/customers", customerHandler.AddCustomer)
	app.Put("/customers/:id", customerHandler.UpdateCustomer)
	app.Delete("/customers/:id", customerHandler.DeleteCustomer)

	err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
	if err != nil {
		log.Fatal(err)
	}

}

func initDatabase() *gorm.DB {

	dial := mysql.Open("root:1234@tcp(host.docker.internal:3306)/customer")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "host.docker.internal:6379",
	})
}
