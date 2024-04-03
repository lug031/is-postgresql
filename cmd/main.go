package main

import (
	"is-postgresql/api/routes"
	"is-postgresql/cmd/server"
	"is-postgresql/cmd/utils/database"
	"is-postgresql/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar enrutador Gin-gonic
	router := gin.Default()

	// Inicializar conexión a la base de datos
	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Middlewares
	router.Use(middleware.GinMiddleware())

	// Configurar rutas
	routes.ClienteRoutes(router)

	// Iniciar servidor
	server.RunServer(router)
}
