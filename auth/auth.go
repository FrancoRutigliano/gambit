package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//Estructura que decodifica el JWT

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_Use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_Id string
	Username  string
}

// Voy a utilizar unicamente Username

// Dividimos el Token, porque este viene en tres partes y se saparan por un punto
// Header
// Payload = contiene la información real que se transmitirá a la aplicación
// Firma
func ValidoToken(token string) (bool, error, string) {
	//En base a un separador dividimos un string en partes
	parts := strings.Split(token, ".")

	//validacion
	if len(parts) != 3 {
		fmt.Println("El token no es válido")
		return false, nil, "El token no es válido"
	}

	// Para que nos devuelva los valores como estan en nuestra estructura
	//decodificamos el payload
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token: ", err.Error())
		return false, nil, err.Error()
	}

	//En el caso de que todo va bien

	var tkj TokenJSON
	//convertimos en una estructura de go, tomamos un slice de bytes y lo convertimos y poder manipular
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar en estructura JSON: ", err.Error())
		return false, err, err.Error()
	}

	//Tambien tengo que validar que el token no este expirado

	ahora := time.Now()
	//Conversion entre Exp y hay que convertirlo en una fecha valida
	tm := time.Unix(int64(tkj.Exp), 0)

	//Yo quiero validar si tm esta antes de la var ahora, quiere decir que esta expirado el token

	if tm.Before(ahora) {
		fmt.Println("Fecha expiración token " + tm.String())
		fmt.Println("Token Expirado!")
		return false, err, "Token expirado !!"
	}

	return true, nil, string(tkj.Username)

}
