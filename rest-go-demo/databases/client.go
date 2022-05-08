package databases

import (
	"log"
	"rest-go-demo/entity"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operations
var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection was succesfully")
	return nil
}

func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
