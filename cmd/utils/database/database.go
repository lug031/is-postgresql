package database

import (
	"errors"
	"fmt"
	"is-postgresql/cmd/utils/log"
	"is-postgresql/pkg/models"
	"os"

	"github.com/joho/godotenv"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB
var logger = log.GetLogger()

func ConnectToDatabase() error {
	logger.Info("Conectando a la base de datos...")

	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error al cargar variables de entorno", zap.Error(err))
		return err
	}

	// Variables de conexion
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePort := os.Getenv("DATABASE_PORT")

	// Verificar si alguna variable de entorno no está definida
	if databaseUser == "" || databasePassword == "" || databaseHost == "" || databaseName == "" || databasePort == "" {
		return errors.New("faltan variables de entorno necesarias para la conexión a la base de datos")
	}

	Connection = connectPostgres(databaseUser, databasePassword, databaseHost, databaseName, databasePort)
	return nil
}

func connectPostgres(databaseUser string, databasePassword string, databaseHost string, databaseName string, databasePort string) *gorm.DB {
	logger.Info(fmt.Sprintf("Conectando a Postgres %s..", databaseHost))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", databaseHost, databaseUser, databasePassword, databaseName, databasePort)
	Connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error("Error al conectar a Postgres", zap.Error(err))
		panic(err)
	}
	err = Connection.AutoMigrate(&models.Cliente{}) // Aquí se agrega la migración automática para el modelo Cliente
	if err != nil {
		logger.Error("Error al realizar la migración automática", zap.Error(err))
		panic(err)
	}

	logger.Info("Conexión exitosa a Postgres")
	return Connection
}

func CheckConnection() bool {
	db, err := Connection.DB()
	if err != nil {
		logger.Error("Error al verificar la conexión a la base de datos", zap.Error(err))
		return false
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Error al verificar la conexión a la base de datos", zap.Error(err))
		return false
	}

	return true
}
