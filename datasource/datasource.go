package datasource

import (
	"fmt"
	"github.com/fabriciolfj/register-product-go/entities"
	"log"
	"os"

	"github.com/magiconair/properties"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDataBase() *gorm.DB {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = p.MustGetString("db_host")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		p.MustGetString("db_user"),
		p.MustGetString("db_password"),
		host,
		p.MustGetString("db_port"),
		p.MustGetString("db_name"),
	)

	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil || DB == nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	err = DB.AutoMigrate(&entities.Product{})

	if err != nil {
		log.Printf("Erro ao realizar a migração: %v", err)
	}

	return DB
}
