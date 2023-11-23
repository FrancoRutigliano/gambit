package bd

import (
	"database/sql"
	"fmt"

	"github.com/FrancoRutigliano/gambit/models"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro de InsertCategory")

	//Concetamos a la Base de datos
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	//Sentencia del Insert SQL

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	//Variable especifica y estructura del paquete sql
	var result sql.Result
	//Db.Exec en Go ejecuta consultas SQL que no devuelven filas, como INSERT, UPDATE o DELETE, en una base de datos.
	result, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	// Tenemos dos valores para result, cantidad de filas afectadas y ult ID insertado.
	// A nosotros nos interesa el segundo, devuelve el ult ID insertado en la tabla
	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}
	//Depuración
	fmt.Println("Insert Category > Ejecución Exitosa")
	return LastInsertId, err2

}
