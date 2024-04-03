package repository

import (
	"is-postgresql/cmd/utils/database"
	"is-postgresql/cmd/utils/log"
	"is-postgresql/pkg/models"

	"go.uber.org/zap"
)

var logger = log.GetLogger()

func GetAllClientes() ([]models.Cliente, error) {
	var clientes []models.Cliente
	result := database.Connection.Find(&clientes)
	if result.Error != nil {
		return nil, result.Error
	}
	return clientes, nil
}

func GetClienteById(id string) (*models.Cliente, error) {
	var cliente models.Cliente
	result := database.Connection.First(&cliente, "dni_cliente = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cliente, nil
}

func AddCliente(cliente *models.Cliente) (*models.Cliente, error) {
	logger.Info("Agregando cliente", zap.Any("cliente", cliente))

	result := database.Connection.Create(cliente)
	if result.Error != nil {
		logger.Error("Error al agregar un cliente", zap.Error(result.Error))
		return nil, result.Error
	}

	logger.Debug("Cliente agregado", zap.Any("cliente", cliente))
	return cliente, nil
}

func UpdateClienteById(id string, cliente *models.Cliente) (*models.Cliente, error) {
	existingCliente, err := GetClienteById(id)
	if err != nil {
		return nil, err
	}

	cliente.DNI_cliente = existingCliente.DNI_cliente

	logger.Info("Actualizando cliente con ID", zap.String("ID", id))
	logger.Debug("Cliente a actualizar", zap.Any("cliente", cliente))

	result := database.Connection.Model(&models.Cliente{}).Where("dni_cliente = ?", id).Updates(cliente)
	if result.Error != nil {
		logger.Error("Error al actualizar el cliente", zap.Error(result.Error))
		return nil, result.Error
	}

	return cliente, nil
}

func DeleteClienteById(id string) error {
	logger.Info("Eliminando cliente con ID", zap.String("ID", id))

	result := database.Connection.Delete(&models.Cliente{}, id)
	if result.Error != nil {
		logger.Error("Error al eliminar el cliente", zap.Error(result.Error))
		return result.Error
	}

	return nil
}
