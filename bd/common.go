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
