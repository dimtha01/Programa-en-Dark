package main

import (
	_"fmt"
	"modulo/util"
)

func main() {
	//Declaracion De Variables
	var (
		numero_1 int = 12
		numero_2 int = 12
		suma     int = 0
	)

	//Entadadas de Datos

	util.Encabezado("Suma", 30)
	numero_1 = util.IngresarDatos("Ingrese El Primer Numero")
	numero_2 = util.IngresarDatos("Ingrese El Segundo Numero")
	suma = util.Suma(numero_1,numero_2)

	//Salida de Datos
	util.Mostrar("Resultado de la Suma es",suma,30)
	

}
