package services

import (
	"errors"
	"is-postgresql/cmd/utils"
	"is-postgresql/cmd/utils/log"
	"is-postgresql/pkg/entities"
	"is-postgresql/pkg/mappers"
	"is-postgresql/pkg/repository"

	"go.uber.org/zap"
)

var logger = log.GetLogger()

func GetAllClientes() []entities.Cliente {
	clientes, err := repository.GetAllClientes()
	if err != nil {
		logger.Error("Error al obtener todos los clientes", zap.Error(err))
		return nil
	}
	mClientes := mappers.ConvertToEntities(clientes)
	return mClientes
}

func GetClienteById(id string) *entities.Cliente {
	cliente, err := repository.GetClienteById(id)
	if err != nil {
		logger.Error("Error al obtener el cliente por ID", zap.Error(err))
		return nil
	}
	mCliente := mappers.ConvertToEntity(cliente)
	return mCliente
}

func AddCliente(cliente *entities.Cliente) (*entities.Cliente, error) {
	logger.Info("Agregando cliente", zap.Any("cliente", cliente))

	edad, err := utils.CalcularEdad(cliente.FechaNacimiento)
	if err != nil {
		return nil, err
	}

	logger.Info("EDAD cliente", zap.Any("edad", edad))

	if edad < 18 {
		return nil, errors.New("solo registrar clientes mayor de 18 años")
	}

	modelCliente := mappers.ConvertToModel(cliente)
	createdModelCliente, err := repository.AddCliente(modelCliente)
	if err != nil {
		logger.Error("Error al agregar un cliente", zap.Error(err))
		return nil, err
	}

	createdEntityCliente := mappers.ConvertToEntity(createdModelCliente)
	logger.Info("Cliente agregado", zap.Any("cliente", createdEntityCliente))
	return createdEntityCliente, nil
}

func UpdateClienteById(id string, cliente *entities.Cliente) (*entities.Cliente, error) {
	logger.Info("Actualizando cliente con ID", zap.String("ID", id))
	logger.Debug("Cliente a actualizar", zap.Any("cliente", cliente))

	updatedCliente, err := repository.UpdateClienteById(id, mappers.ConvertToModel(cliente))
	if err != nil {
		logger.Error("Error al actualizar el cliente", zap.Error(err))
		return nil, err
	}

	mCliente := mappers.ConvertToEntity(updatedCliente)
	return mCliente, nil
}

func DeleteClienteById(id string) error {
	logger.Info("Eliminando cliente con ID", zap.String("ID", id))

	cliente, err := repository.GetClienteById(id)
	if err != nil {
		logger.Error("Error al obtener el cliente por ID", zap.Error(err))
		return err
	}

	edad, err := utils.CalcularEdad(cliente.FechaNac_cliente)
	if err != nil {
		return err
	}

	logger.Info("EDAD cliente", zap.Any("edad", edad))

	if edad < 80 {
		return errors.New("solo se pueden eliminar clientes mayores de 80 años")
	}

	err = repository.DeleteClienteById(id)
	if err != nil {
		logger.Error("Error al eliminar el cliente", zap.Error(err))
		return err
	}
	return nil
}
