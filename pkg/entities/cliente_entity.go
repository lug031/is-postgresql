package entities

type Cliente struct {
	Dni             uint64 `json:"dni"`
	Nombres         string `json:"nombres"`
	Apellidos       string `json:"apellidos"`
	FechaNacimiento string `json:"fechaNacimiento"`
	Sexo            string `json:"sexo"`
	Ciudad          string `json:"ciudad"`
}
