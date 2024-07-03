package util

import (
	"fmt"
	"strings"
)

func Linea(largo int) {
	fmt.Println(strings.Repeat("═", largo))
}

func Encabezado(titulo string, largo int) {
	Linea(largo)
	fmt.Println("\t", strings.ToUpper(titulo))
	Linea(largo)
}
func Suma(numero_1 int, numero_2 int) int {

	return numero_1 + numero_2

}
func IngresarDatos(msj string) int {
	var numero int = 0
	fmt.Printf("%v : ", msj)
	fmt.Scanln(&numero)
	return numero

}
func Mostrar(msj string, resultado int, largo int) {
	Linea(largo)
	println(msj, " : ", resultado)
	Linea(largo)
}
func LeerText(msj string) string {
	texto := ""
	fmt.Printf("%v : ", msj)
	fmt.Scanln(&texto)
	return texto
}
func LeerEntero(msj string) int {
	var numeroLeido int = 0
	fmt.Printf("%v : ",msj)
	fmt.Scanln(&numeroLeido)

	return numeroLeido
}
