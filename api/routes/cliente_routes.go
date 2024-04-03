package routes

import (
	"is-postgresql/api/controllers"
	"net/http"

	"is-postgresql/pkg/entities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validateRequest(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}

func ClienteRoutes(router *gin.Engine) {
	router.GET("/clientes", controllers.GetAllClientes)
	router.GET("/clientes/:id", controllers.GetClienteById)
	router.POST("/clientes", func(c *gin.Context) {
		var req entities.Cliente
		if err := validateRequest(c, &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		controllers.AddCliente(c, &req)
	})
	router.PUT("/clientes/:id", func(c *gin.Context) {
		var req entities.Cliente
		if err := validateRequest(c, &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		controllers.UpdateClienteById(c, &req)
	})
	router.DELETE("/clientes/:id", controllers.DeleteClienteById)
}
