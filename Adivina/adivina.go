package main

import (
	"fmt"
	"modulo/util"
	"strings"
)

func main() {
	var (
		alias         string = ""
		numero        int    = 0
		numeroSecreto int    = 12
	)

	util.Encabezado("Juego de Adivina", 35)
	alias = util.LeerText("Ingresa tu Alias")
	jugar(alias,numero,numeroSecreto)
}

// Funciones deL Programa Adivina
func jugar(alias string, numero int, numeroSecreto int){
	for numero != numeroSecreto{
		util.Linea(35)
		numero = util.LeerEntero("Ingrese el Numero ")
		fmt.Println("")
		util.Linea(35)
		if numero > numeroSecreto {
			util.Linea(35)
			fmt.Println("El numero es mayor que el secreto " )
			util.Linea(35)
		}else if numero < numeroSecreto {
			util.Linea(35)
			fmt.Println("El numero es menor que el secreto ")
			util.Linea(35)
		}else{
			util.Linea(35)
			fmt.Println("Ganastes te Felicidades " + strings.ToUpper(alias))
			util.Linea(35)
		}
	}
}
