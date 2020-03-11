package main

// Examen Julia Elizabeth Garcia Herrera ITICS 10-1

import (
	"bufio"                                                
	"fmt"                              
	"os" 
	"io/ioutil"                       
)

type Alumno struct {
	Nombre, Calificacion string
}


func main() {
	menu := `¿Qué quieres hacer?
[1] -- Capturar
[2] -- Mostrar
[3] -- Salir

----->	`
	var eleccion int
	
	for eleccion != 3 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			var Alumnos []Alumno
			for i:= 0; i<5; i++ {
			var c Alumno
			fmt.Println("Ingresa el nombre:")
			if scanner.Scan() {
				c.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa la Calificacion:")
			if scanner.Scan() {
				c.Calificacion = scanner.Text()
			}
		
			Alumnos = append(Alumnos,c)
		}
			
		for i:= 0; i<len(Alumnos); i++ {
			InsertarDatos(Alumnos[i].Nombre, Alumnos[i].Calificacion)
		}
		
		case 2:
			
			data, err := ioutil.ReadFile("text.txt")
    		if err != nil {
        		fmt.Println(err)
    		}

    		// Should output: `The file contains: Hello, world!`
    		fmt.Println("El archivo contiene \n " + string(data))

			
		
			
			
		}
	}
}
func InsertarDatos (Nombre string, Calificacion string) {
	f, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(Nombre+" "+Calificacion+"\n"); err != nil {
		fmt.Println(err)
	}

}
