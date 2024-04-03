package utils

import (
	"errors"
	"time"
)

func CalcularEdad(fechaNacimiento string) (int, error) {
	const formatoFecha = "02/01/2006"
	fechaNac, err := time.Parse(formatoFecha, fechaNacimiento)
	if err != nil {
		return 0, errors.New("formato de fecha de nacimiento inválido")
	}

	hoy := time.Now()

	edad := hoy.Year() - fechaNac.Year()
	if hoy.Month() < fechaNac.Month() || (hoy.Month() == fechaNac.Month() && hoy.Day() < fechaNac.Day()) {
		edad--
	}

	return edad, nil
}
