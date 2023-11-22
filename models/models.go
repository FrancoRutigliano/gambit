package models

/*
Vamos a asignarles valores a nuestras variables del struct, valores que vienen desde el json y se lo indicamos desde que campo con los: ``
*/

type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

//Estructura que va a contener los datos del registro

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
