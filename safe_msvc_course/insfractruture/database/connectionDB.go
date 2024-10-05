package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/safe_msvc_course/insfractruture/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345678"
	DB_NAME     = "msvc_safe_course_db"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_SSLMODE  = "disable"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println(errEnv.Error())
	}

	strConnection := CreateDatabase()

	dsn := fmt.Sprintf(strConnection+" dbname=%s", DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	MigrateDatabase(db)
	return db
}
func CreateDatabase() string {
	strConnection := fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_SSLMODE)
	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	query := fmt.Sprintf("SELECT 1 FROM  pg_database WHERE datname ='%s'", DB_NAME)
	errd := db.Exec(query)
	if errd.RowsAffected == 0 {
		// Crear la base de datos usando una consulta SQL cruda
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s", DB_NAME)
		err = db.Exec(createDBSQL).Error
		if err != nil {
			log.Println(err.Error())
		}
		log.Printf("Base de datos '%s' creada exitosamente.\n", DB_NAME)
	}
	return strConnection
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&entities.Course{}, &entities.Topic{}, &entities.Language{})
}
func CloseConnection() {
	var db *gorm.DB = DatabaseConnection()
	dbSQL, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	dbSQL.Close()
}
