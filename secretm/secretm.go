package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/FrancoRutigliano/gambit/awsgo"
	"github.com/FrancoRutigliano/gambit/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto " + nombreSecret)

	//Inicializacion a secretM mediante la variable Cfg
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	//Unmarshal parsea el json codificado y lo convierte y envia a la structura
	// Y así estamos procesando lo que devuelve el secretvalue
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret OK " + nombreSecret)

	return datosSecret, nil
}
