package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PaymentPromise struct {
	gorm.Model
	Bill     string
	Document string
	Client   string
	Date     string
}

func createPaymentPromise(paymentPromise map[string]interface{}) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	dateStr := paymentPromise["date"].(string)
	date, err := time.Parse("02/01/2006", dateStr)

	if err != nil {
		log.Println(err)
	}

	formattedDate := date.Format("2006-01-02")
	pp := &PaymentPromise{
		Bill:     paymentPromise["bill"].(string),
		Document: paymentPromise["document"].(string),
		Client:   paymentPromise["client"].(string),
		Date:     formattedDate,
	}

	result := db.Create(pp)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
