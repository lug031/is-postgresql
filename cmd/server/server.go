package server

import (
	"is-postgresql/cmd/utils/database"
	logger "is-postgresql/cmd/utils/log"

	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
	"go.uber.org/zap"
)

func RunServer(router *gin.Engine) {
	log := logger.GetLogger()

	// Crea un manejador de healthcheck
	health := healthcheck.NewHandler()

	health.AddLivenessCheck("check server", func() error {
		return nil
	})

	health.AddReadinessCheck("database", func() error {
		if !database.CheckConnection() {
			log.Error("Error al verificar la conexión a la base de datos")
		}
		return nil
	})

	// Registra el path /liveness y /readiness para que se puedan hacer pruebas de salud
	router.GET("/liveness", gin.WrapF(health.LiveEndpoint))
	router.GET("/readiness", gin.WrapF(health.ReadyEndpoint))

	// Iniciar el servidor
	err := router.Run(":8080")
	if err != nil {
		log.Error("Error al iniciar el servidor", zap.Error(err))
	}
}
