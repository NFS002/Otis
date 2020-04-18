package client

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgresql driver
	"log"
	"os"
)



// RDSClient : Struct to represent a connection to a DynamoDB instance
type RDSClient struct {
	DB *gorm.DB
}

// Init : Function called on startup to initialize the amazon RDS connection
func (c *RDSClient) Init() error {
	var name = os.Getenv("DB_NAME")
	var port = os.Getenv("DB_PORT")
	var host = os.Getenv("DB_ENDPOINT")
	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASSWORD")
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, pass)
	fmt.Println(connectionStr)
	db, err := gorm.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	pingErr := db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		return pingErr
	}
	c.DB = db
	return nil
}