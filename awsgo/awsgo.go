package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

//var de tipo context

var Ctx context.Context

// var de tipo config, estructura que tiene todo lo necesario para manejar las configuraciones de inicio de sesion
var Cfg aws.Config
var err error

func InicializoAWS() {
	// Le indicamos que no hay variables de entorno, contextos limitantes. Es contexto vacio
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	//Tratamiento de error

	if err != nil {
		panic("Error al cargar la configurations .aws/ config" + err.Error())
	}
}
