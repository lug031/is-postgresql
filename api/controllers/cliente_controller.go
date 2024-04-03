package controllers

import (
	"is-postgresql/cmd/utils/log"
	"is-postgresql/pkg/entities"
	"is-postgresql/pkg/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger = log.GetLogger()

func GetAllClientes(c *gin.Context) {
	logger.Debug("Inicio de método GetAllClientes")
	clientes := services.GetAllClientes()
	c.JSON(200, clientes)
}

func GetClienteById(c *gin.Context) {
	id := c.Param("id")
	logger.Debug("GetClienteById - ID a buscar", zap.String("ID", id))

	cliente := services.GetClienteById(id)
	if cliente == nil {
		logger.Warn("Registro con ID no encontrado", zap.String("ID", id))
		c.JSON(404, gin.H{"message": "No se encontró el registro"})
		return
	}

	c.JSON(200, cliente)
}

func AddCliente(c *gin.Context, cliente *entities.Cliente) {
	createdCliente, err := services.AddCliente(cliente)
	if err != nil {
		logger.Error("Error al agregar un cliente", zap.Error(err))
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	logger.Info("Cliente agregado", zap.Any("cliente", createdCliente))
	c.JSON(201, gin.H{"message": "Cliente agregado", "cliente": createdCliente})
}

func UpdateClienteById(c *gin.Context, cliente *entities.Cliente) {
	id := c.Param("id")

	updatedCliente, err := services.UpdateClienteById(id, cliente)
	if err != nil {
		logger.Error("Error al actualizar el cliente", zap.Error(err))
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	logger.Info("Cliente actualizado", zap.Any("cliente", updatedCliente))
	c.JSON(200, gin.H{"message": "Cliente actualizado", "cliente": updatedCliente})
}

func DeleteClienteById(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteClienteById(id)
	if err != nil {
		logger.Error("Error al eliminar el cliente", zap.Error(err))
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	logger.Info("Cliente eliminado", zap.String("ID", id))
	c.JSON(200, gin.H{"message": "Cliente eliminado"})
}
