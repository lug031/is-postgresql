package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/rs/cors"
    "github.com/sirupsen/logrus"
    "os"
)

func GinMiddleware() gin.HandlerFunc {
    // Configurar el middleware de CORS
    corsConfig := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    })

    // Configurar el middleware de registro
    logFile, err := os.Create("server.log")
    if err != nil {
        logrus.Fatal("No se pudo crear el archivo de registro:", err)
    }
    gin.DefaultWriter = logFile

    return func(c *gin.Context) {
        // Aplicar el middleware de CORS
        corsConfig.HandlerFunc(c.Writer, c.Request)

        // Ejecutar el siguiente middleware
        c.Next()
    }
}
