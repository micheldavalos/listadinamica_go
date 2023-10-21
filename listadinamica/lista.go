package listadinamica

import (
	"errors"
)

type ListaDinamica struct {
	data []int64
}

func (l *ListaDinamica) InsertarFinal(dato int64) {
	l.data = append(l.data, dato)
}

func (l *ListaDinamica) InsertarInicio(dato int64) {
	nuevo := []int64{dato}
	l.data = append(nuevo, l.data...)
}

func (l *ListaDinamica) InsertarPosicion(dato int64, pos int) error {
	if pos < 0 || pos >= len(l.data) {
		return errors.New("La posición no es válida")
	}

	primerParte := l.data[:pos]
	segundaParte := l.data[pos:]

	var nuevo []int64
	nuevo = append(nuevo, primerParte...)
	nuevo = append(nuevo, dato)
	nuevo = append(nuevo, segundaParte...)

	l.data = nuevo

	return nil
}

func (l *ListaDinamica) EliminarFinal() error {
	if len(l.data) == 0 {
		return errors.New("La lista está vacía")
	}

	l.data = l.data[:len(l.data)-1]
	return nil
}

func (l *ListaDinamica) EliminarInicio() error {
	if len(l.data) == 0 {
		return errors.New("La lista está vacía")
	}

	l.data = l.data[1:]

	return nil
}

func (l *ListaDinamica) EliminarPosicion(pos int) error {
	if pos < 0 || pos >= len(l.data) {
		return errors.New("Posición no válida")
	}
	// l.data[pos+1:] = [3, 4, 5]
	l.data = append(l.data[:pos], l.data[pos+1:]...)

	return nil
}

