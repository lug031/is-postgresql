package models

import (
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Cliente struct {
	DNI_cliente       uint64 `gorm:"primaryKey;column:dni_cliente"`
	Nombres_cliente   string `gorm:"varchar(255);not null;column:nombres_cliente"`
	Apellidos_cliente string `gorm:"varchar(255);not null;column:apellidos_cliente"`
	FechaNac_cliente  string `gorm:"varchar(10);not null;column:fecha_nac_cliente"`
	Sexo_cliente      string `gorm:"varchar(1);not null;column:sexo_cliente"`
	Ciudad_cliente    string `gorm:"varchar(255);not null;column:ciudad_cliente"`
}

func (c *Cliente) BeforeCreate(tx *gorm.DB) (err error) {
	if err := validarDNICliente(c.DNI_cliente); err != nil {
		return err
	}

	if err := validarFormatoFecha(c.FechaNac_cliente); err != nil {
		return err
	}

	if err := validarSexoCliente(c.Sexo_cliente); err != nil {
		return err
	}

	if err := validarErrorGen(); err != nil {
		return err
	}

	return nil
}

func validarDNICliente(dni uint64) error {
	dniStr := strconv.FormatUint(dni, 10)
	if len(dniStr) != 8 {
		return errors.New("el DNI debe tener exactamente 8 dígitos")
	}
	return nil
}

func validarFormatoFecha(fecha string) error {
	parts := strings.Split(fecha, "/")
	if len(parts) != 3 {
		return errors.New("el formato de la fecha de nacimiento debe ser dd/mm/yyyy")
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("el día de la fecha de nacimiento no es válido")
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return errors.New("el mes de la fecha de nacimiento no es válido")
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return errors.New("el año de la fecha de nacimiento no es válido")
	}

	if day < 1 || day > 31 {
		return errors.New("el día de la fecha de nacimiento debe estar entre 1 y 31")
	}
	if month < 1 || month > 12 {
		return errors.New("el mes de la fecha de nacimiento debe estar entre 1 y 12")
	}
	if year < 1900 {
		return errors.New("el año de la fecha de nacimiento no es válido")
	}

	return nil
}

func validarSexoCliente(sexo string) error {
	if sexo != "M" && sexo != "F" {
		return errors.New("el sexo debe ser 'M' o 'F'")
	}
	return nil
}

func validarErrorGen() error {
	return nil
}
