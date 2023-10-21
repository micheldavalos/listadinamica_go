package tests

import (
	"testing"
)

func TestArregloFijo(t *testing.T) {
	t.Log("Test arreglo fijo")

	arreglo := [5]int64{1, 2, 3, 4, 5}

	t.Logf("arreglo: %v", arreglo)
	t.Logf("cap: %d", cap(arreglo))
	t.Logf("len: %d", len(arreglo))

}

func TestArregloFijo2(t *testing.T) {
	t.Log("Test arreglo fijo")

	var arreglo [5]int64

	arreglo[2] = 10

	t.Logf("arreglo: %v", arreglo)
	t.Logf("cap: %d", cap(arreglo))
	t.Logf("len: %d", len(arreglo))
}

func TestSlice(t *testing.T) {
	t.Log("Test slices")

	var slice []int64

	t.Logf("slice: %v", slice)
	t.Logf("cap: %d", cap(slice))
	t.Logf("len: %d", len(slice))

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7)
	t.Logf("slice: %v", slice)
	t.Logf("cap: %d", cap(slice))
	t.Logf("len: %d", len(slice))

}

func TestSlice2(t *testing.T) {
	t.Log("Test slice con make")

	slice := make([]int64, 100)

	t.Logf("slice: %v", slice)
	t.Logf("cap: %d", cap(slice))
	t.Logf("len: %d", len(slice))
}

func TestSlicing(t *testing.T) {
	t.Log("Test slicing")

	arreglo := [10]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	slicing := arreglo[1:5]
	arreglo[3] = 500
	
	t.Logf("arreglo %v", arreglo)
	t.Logf("slicing %v", slicing)
	t.Logf("arreglo cap: %d", cap(arreglo))
	t.Logf("arreglo len: %d", len(arreglo))
	t.Logf("slicing cap: %d", cap(slicing))
	t.Logf("slicing len: %d", len(slicing))
}

func TestSlicingCopia(t *testing.T) {
	arreglo := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := make([]int64, len(arreglo))

	copy(slice, arreglo)

	t.Logf("arreglo %v", arreglo)
	t.Logf("slice %v", slice)

}