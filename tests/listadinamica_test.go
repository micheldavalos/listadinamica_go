package tests

import  (
	"testing"
	"listadinamica/listadinamica"
)


func TestInsertarFinal(t *testing.T) {
	t.Log("Test Insertar al Final")

	ld := listadinamica.ListaDinamica{}
	ld.InsertarFinal(10)
	ld.InsertarFinal(1)
	ld.InsertarFinal(15)
}

func TestInsertarInicio(t *testing.T) {
	t.Log("Test Insertar al Inico")

	ld := listadinamica.ListaDinamica{}
	ld.InsertarFinal(10)
	ld.InsertarFinal(1)
	ld.InsertarFinal(15)
	ld.InsertarInicio(2)
}

func TestInsertarPosicion(t *testing.T) {
	t.Log("Test Insertar en Posición")

	ld := listadinamica.ListaDinamica{}
	ld.InsertarFinal(10)
	ld.InsertarFinal(1)
	ld.InsertarFinal(15)
	ld.InsertarInicio(2)

	err := ld.InsertarPosicion(500, 2)
	if err != nil {
		t.Error(err)
	}
	t.Logf("lista: %v", ld)
}

func TestEliminarFinal(t *testing.T) {
	t.Log("Test Eliminar al Final")

	ld := listadinamica.ListaDinamica{}
	ld.InsertarFinal(10)
	ld.InsertarFinal(1)
	ld.InsertarFinal(15)
	ld.InsertarInicio(2)

	t.Logf("lista: %v", ld)
	err := ld.EliminarFinal()
	if err != nil {
		t.Error(err)
	}
	t.Logf("lista: %v", ld)
}

func TestEliminarInicio(t *testing.T) {
	t.Log("Test Eliminar al Inicio")

	ld := listadinamica.ListaDinamica{}
	ld.InsertarFinal(10)
	ld.InsertarFinal(1)
	ld.InsertarFinal(15)
	ld.InsertarInicio(2)

	t.Logf("lista: %v", ld)
	err := ld.EliminarInicio()
	if err != nil {
		t.Error(err)
	}
	t.Logf("lista: %v", ld)
}

func TestEliminarPosicion(t *testing.T) {
	t.Log("Test Eliminar por Posición")

	ld := listadinamica.ListaDinamica{}
	ld.InsertarFinal(10)
	ld.InsertarFinal(1)
	ld.InsertarFinal(15)
	ld.InsertarInicio(2)

	t.Logf("lista: %v", ld)
	err := ld.EliminarPosicion(2)
	if err != nil {
		t.Error(err)
	}
	t.Logf("lista: %v", ld)
}