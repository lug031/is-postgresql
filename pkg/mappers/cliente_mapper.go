package mappers

import (
	"is-postgresql/pkg/entities"
	"is-postgresql/pkg/models"
)

func ConvertToModels(clientes []entities.Cliente) []models.Cliente {
	var convertedClientes []models.Cliente
	for _, cliente := range clientes {
		convertedCliente := models.Cliente{
			DNI_cliente:       cliente.Dni,
			Nombres_cliente:   cliente.Nombres,
			Apellidos_cliente: cliente.Apellidos,
			FechaNac_cliente:  cliente.FechaNacimiento,
			Sexo_cliente:      cliente.Sexo,
			Ciudad_cliente:    cliente.Ciudad,
		}
		convertedClientes = append(convertedClientes, convertedCliente)
	}
	return convertedClientes
}

func ConvertToEntities(clientes []models.Cliente) []entities.Cliente {
	var convertedClientes []entities.Cliente
	for _, cliente := range clientes {
		convertedCliente := entities.Cliente{
			Dni:             cliente.DNI_cliente,
			Nombres:         cliente.Nombres_cliente,
			Apellidos:       cliente.Apellidos_cliente,
			FechaNacimiento: cliente.FechaNac_cliente,
			Sexo:            cliente.Sexo_cliente,
			Ciudad:          cliente.Ciudad_cliente,
		}
		convertedClientes = append(convertedClientes, convertedCliente)
	}
	return convertedClientes
}

func ConvertToEntity(cliente *models.Cliente) *entities.Cliente {
	convertedCliente := entities.Cliente{
		Dni:             cliente.DNI_cliente,
		Nombres:         cliente.Nombres_cliente,
		Apellidos:       cliente.Apellidos_cliente,
		FechaNacimiento: cliente.FechaNac_cliente,
		Sexo:            cliente.Sexo_cliente,
		Ciudad:          cliente.Ciudad_cliente,
	}
	return &convertedCliente
}

func ConvertToModel(cliente *entities.Cliente) *models.Cliente {
	convertedCliente := models.Cliente{
		DNI_cliente:       cliente.Dni,
		Nombres_cliente:   cliente.Nombres,
		Apellidos_cliente: cliente.Apellidos,
		FechaNac_cliente:  cliente.FechaNacimiento,
		Sexo_cliente:      cliente.Sexo,
		Ciudad_cliente:    cliente.Ciudad,
	}
	return &convertedCliente
}
