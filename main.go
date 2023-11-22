package main

import (
	"context"
	"os"
	"strings"

	"github.com/FrancoRutigliano/gambit/awsgo"
	"github.com/FrancoRutigliano/gambit/bd"
	"github.com/FrancoRutigliano/gambit/handlers"
	"github.com/aws/aws-lambda-go/events"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//Con esta funcion vamos a "arrancar"
	lambda.Start(EjecutoLambda)
}

/*
ctx context.Context: El parámetro ctx se refiere a un contexto de ejecución. En Go (el lenguaje de programación en el que está escrita esta función), el contexto se utiliza para controlar y gestionar la cancelación, el tiempo de espera y otros aspectos de la ejecución de una función.

request events.APIGatewayV2HTTPRequest: El parámetro request es un objeto que representa la solicitud HTTP que ha sido recibida por la función a través de API Gateway. En este caso, events.APIGatewayV2HTTPRequest es un tipo de estructura proporcionada por el AWS Lambda Go SDK que contiene información sobre la solicitud HTTP entrante
*/
func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	// Funcion que se encarga de inicializar y cargar configuraciones de aws
	awsgo.InicializoAWS()
	//Esto sucede en caso de no tener los parametros que necesito
	if !ValidoParametros() {
		panic("Error en los parametros debe enviar 'SecretName', 'UrlPrefix'")
	}

	//var de respuesta, donde se configura todo lo que debamos devolver de respuesta.
	var res *events.APIGatewayProxyResponse
	// deve devolver la ruta - el prefijo 'prefix'
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	// Llamado al Handler
	status, message := handlers.Manejadores(path, method, body, header, request)

	headerResp := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headerResp,
	}

	return res, nil

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}
	return traeParametro
}
