package multimedia

import "fmt"

// Creando structura con un slice para almacenar contenido multimedia
type ContenidoWeb struct {
	Contenido []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	for _, c := range cw.Contenido {
		fmt.Println(c.Mostrar())
	}
}

// INTERFACE _______________________________________________________________________
type Multimedia interface {
	Mostrar() (string, string, int64)
}

// ESTRUCTURAS Y MÉTODOS ___________________________________________________________
type Imagen struct {
	Titulo  string
	Formato string
	Canales int64
}

func (i *Imagen) Mostrar() (string, string, int64) {
	return i.Titulo, i.Formato, i.Canales
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int64
}

func (a *Audio) Mostrar() (string, string, int64) {
	return a.Titulo, a.Formato, a.Duracion
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (v *Video) Mostrar() (string, string, int64) {
	return v.Titulo, v.Formato, v.Frames
}
