package tools

import (
	"fmt"
	"time"
)

//En MySql las fechas se alamcenan en un formato predeterminado, Año - Mes - Dia t\ H:M:S

func FechaMySQL() string {
	// Variable fecha de hoy, sin formatear
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
