/*
Ultimo avance
Walther Ramirez
29 enero 2019 
*/
package main

import (          
		"fmt"
		"errors"
		"math"
)

type persona struct {  // CREAR ESTRUCTURAS FUERA DEL MAIN 
	nombre string
	edad int 
}

func main() {
	fmt.Println("Estructura con el prefijo de type modo herencia")
	p := persona{nombre:"Aaron",edad:27}
	fmt.Println(p.nombre, p.edad)
	fmt.Println("========================================")
	fmt.Println("Punteros")
	w := "Walther"
	fmt.Println(&w, "Es el valor del puntero")
	
	resultado := sumar(4,5)
	fmt.Println(resultado, "Este es el resultado de otra funcion matematica!") // IMPRIMIMOS EL VALOR DE LA OTRA FUNCION FUERA DEL MAIN 
	fmt.Println("========================================")
	fmt.Println("Funciones con el uso de errores error Generado cuando el numero es negativo")
	resultado2, err := sqrt(-1)
	if err != nil{
		fmt.Println(err)
		}else{
		fmt.Println(resultado2)       // IMPRIMIMOS EL VALOR DE LA FUNCION  sqrt FUERA DEL MAIN
		}
	
	
	
	fmt.Println("========================================")
	fmt.Println("Prueba de arreglos")
	
	var a [5]int // Arrays
	a[0]= 26     // seteaer en el arreglo 
	fmt.Println(a, "Arreglo A")       //fmt.Println imprimir
	fmt.Println("========================================")
	fmt.Println("Prueba de arreglos concatenados")
	b := [] int {1,2,3}
	//b = append(a,13)
	fmt.Println(b,"Arreglo B")
	fmt.Println("========================================")
	x := 1
	y := 22
	suma := x + y
	fmt.Println("La suma es de ",suma)
	fmt.Println("========================================")
	fmt.Println("Manejo de condiciones")
	// MANEJO DE CONDICIONES 
	fmt.Println("Condicion para los valores de X y Y")
	if x > y {				//Manera para condicionar 
		fmt.Println("X es mayor que y")
	} else {
		fmt.Println("y es mayor que x")	
	}
	fmt.Println("========================================")
	
	Figuras := make(map[string]int) // concatenar funciones dentro de un arreglo 
	Figuras["Triangulo"]= 3
	Figuras["Cuadrado"]= 4
	Figuras["Circulo"]= 0
	delete(Figuras, "Circulo") //Eliminar parametros 
    fmt.Println(Figuras)
	fmt.Println(Figuras["Cuadrado"]) // De esta manera se puede asignar a un solo valor
	
	// MANEJO DE LOOPS 
	for i := 0; i < 5; i++ {
		fmt.Println("Repite ", i)
	}
	fmt.Println("========================================")
	fmt.Println("Cambio de uso de loops concatenados")
	m := make(map[string]string)
	m["1ero"] = "Go"
	m["2do"] = "Javascript"
    m["3ero"] = "Angular"
	
	for key, valor := range m{
		fmt.Println("key", key,"valor", valor )
	}
	
}

func sumar(x int ,y int) int { // CREAMOS OTRA FUNCION FUERA DE LA MAIN CON VALORES NUMERICOS 
		return x + y
	}
	
	
func sqrt( x float64)(float64, error){  // FUNCION PARA 2 VARIABLES Y CON VALORES DE ERROR 
	if x < 0 {
		return 0,errors.New("El valor es negativo")
	}
	return math.Sqrt(x), nil 
	}
	
