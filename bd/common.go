package bd

// funciones comunes  de bases de datos

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/FrancoRutigliano/gambit/models"
	"github.com/FrancoRutigliano/gambit/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error

// Todo lo que sea base de datos se maneja con punteros, por temas de velocidad
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//Ping check que esta todo bien en cuanto a la conexion con db
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	//Si es para depurar sirve, en produ no conviene
	fmt.Println(dsn)

	return dsn
}

/*
	Funcion con la tarea de indicar si es administrador o no
*/

func UserIsAdmin(UserUUID string) (bool, string) {
	//Depuracion
	fmt.Println("Comienza UserIsAdmin")
	//Nos conectamos a la base de dato mediante el objeto global DB
	err := DbConnect()
	//Si hubo algun error se asume que el usuario no es admin y se cierra la DB
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + UserUUID + "' AND User_Status = 0"
	fmt.Println(sentencia)

	//Utilizamos Query para generar una consulta hacia la base de datos en base a la sentencia escrita arriba
	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}
	//Caso de que la sentencia sea correcta, ahora a validar si realmente es admin

	var valor string
	//Posicionamos lo que nos devolvio la sentencia en el primer registro
	rows.Next()
	// Estamos extrayendo los valores con el scan y almacenandolos/destino es la VAR valor
	rows.Scan(&valor)

	fmt.Println("UserIsAdmin > Ejecución Existosa - valor devuelto -> " + valor)
	//Comprobamos de que realmente sea usuario
	if valor == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}
