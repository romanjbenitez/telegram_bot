package Config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"telegram_bot/Models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "q"
	dbname   = "finances"
)

func GetConnection() *gorm.DB {
	dns := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panic("fallo al conectar a la base de datos")
	}
	err = db.AutoMigrate(&Models.Transaccion{})
	if err != nil {
		log.Panic("fallaron las migraciones")
	}

	log.Println("Db established")

	return db
}
