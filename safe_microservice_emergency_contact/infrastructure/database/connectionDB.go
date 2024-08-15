package database

import (
	"fmt"
	"log"
	"os"

	"github.com/all_is_safe/infrastructure/entities"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv.Error())
	}
	DB_NAME := os.Getenv("DB_NAME")

	strConnection := CreateDatabase(DB_NAME)
	dsn := fmt.Sprintf(strConnection+"  dbname=%s",
		DB_NAME,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&entities.EmergencyContact{},)
	return db
}

// Crear la nueva base de datos
func CreateDatabase(DB_NAME string) string {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	strConnection := fmt.Sprintf("host=%s user=%s password=%s  port=%s  sslmode=%s ",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_PORT,
		DB_SSLMODE,
	)

	// Conectar al servidor PostgreSQL sin especificar una base de datos
	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s'", DB_NAME)
	errd := db.Exec(query)
	if errd.RowsAffected == 0 {
		// Crear la base de datos usando una consulta SQL cruda
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s", DB_NAME)
		err = db.Exec(createDBSQL).Error

		if err != nil {
			log.Fatalf("Error al crear la base de datos: %v", err)
		}

		fmt.Printf("Base de datos '%s' creada exitosamente.\n", DB_NAME)
	}

	return strConnection
}
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	dbSQL.Close()
}

func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)
}
