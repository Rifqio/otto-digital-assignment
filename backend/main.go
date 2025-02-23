package main

import (
	"log"
	"net/http"
	"os"
	"voucher-app/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/vouchers?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to the database!")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	routes := routes.MakeRouter(db)
	routes.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port
	}
	e.Logger.Fatal(e.Start(":" + port))
}
