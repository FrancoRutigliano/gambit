package routers

import (
	"encoding/json"
	"strconv"

	"github.com/FrancoRutigliano/gambit/bd"
	"github.com/FrancoRutigliano/gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	//creando instancia de la estructura
	var t models.Category

	//Primero tenemos que decodificar el body que es el que me trae los campos para que yo los pueda validar de manera correcta
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}
	//Validacion para que podamos verificar si el campo CategName también esta correcto
	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la categoria"
	}
	//Lo mismo para path
	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la categoria"
	}

	/*
		Debemos validar si el usuario es admin o no.
	*/
	isAdmin, msg := bd.UserIsAdmin(User)
	// Primera validación
	if !isAdmin {
		return 400, msg
	}
	//En result quiero almacenar el id de la categoria
	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la categoria " + t.CategName + " > " + err2.Error()
	}
	//devolvemos en formato json para que luego pueda ser tomado por el front
	return 200, "{CategID:" + strconv.Itoa(int(result)) + "}"

}
