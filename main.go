package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	var opc int

	cw := multimedia.ContenidoWeb{Contenido: []multimedia.Multimedia{}}

	for true {
		fmt.Println("Selecciona una opción: ")
		fmt.Println("1. Capturar Imagen")
		fmt.Println("2. Capturar Audio")
		fmt.Println("3. Capturar Video")
		fmt.Println("4. Mostrar")
		fmt.Println("5. Salir")

		fmt.Scan(&opc)

		switch {
		// IMAGEN _______________________________________________________________
		case opc == 1:
			var titulo string
			var formato string
			var canales int64

			fmt.Println("CAPTURAR IMAGEN")
			fmt.Println("Título: ")
			fmt.Scan(&titulo)

			fmt.Println("Formato: ")
			fmt.Scan(&formato)

			fmt.Println("Canales: ")
			fmt.Scan(&canales)

			img := multimedia.Imagen{Titulo: titulo, Formato: formato, Canales: canales}

			cw = multimedia.ContenidoWeb{
				Contenido: []multimedia.Multimedia{
					&img,
				},
			}

			// contenido := multimedia.ContenidoWeb{} //inicializacion

		// AUDIO ________________________________________________________________
		case opc == 2:
			var titulo string
			var formato string
			var duracion int64

			fmt.Println("CAPTURAR AUDIO")
			fmt.Println("Título: ")
			fmt.Scan(&titulo)

			fmt.Println("Formato: ")
			fmt.Scan(&formato)

			fmt.Println("Duración: ")
			fmt.Scan(&duracion)

			a := multimedia.Audio{Titulo: titulo, Formato: formato, Duracion: duracion}

			cw = multimedia.ContenidoWeb{
				Contenido: []multimedia.Multimedia{
					&a,
				},
			}

		// VIDEO ________________________________________________________________
		case opc == 3:
			var titulo string
			var formato string
			var frames int64

			fmt.Println("CAPTURAR VIDEO")
			fmt.Println("Título: ")
			fmt.Scan(&titulo)

			fmt.Println("Formato: ")
			fmt.Scan(&formato)

			fmt.Println("Frames: ")
			fmt.Scan(&frames)

			v := multimedia.Video{Titulo: titulo, Formato: formato, Frames: frames}

			cw = multimedia.ContenidoWeb{
				Contenido: []multimedia.Multimedia{
					&v,
				},
			}

		// Mostrar _______________________________________________________________
		case opc == 4:
			cw.Mostrar()

		// Salir _______________________________________________________________
		case opc == 5:
			return
		}
	}
}
