package mem

import (
	"testing"
)

type SimpleStruct struct {
	A float64
	B int
	C string
}

func TestSimpleStruct(t *testing.T) {
	testStruct := SimpleStruct{
		A: 1.,
		B: 2,
		C: "3",
	}
	actualSize := GetMemoryConsumption(testStruct)
	expectedSize := 48
	if actualSize != expectedSize {
		t.Errorf("Size of test struct is incorrect (got %v, expected %v)\n", actualSize, expectedSize)
	}
}

type ComplexStruct struct {
	A float64
	B int
	C string
	D SimpleStruct
	E *SimpleStruct
	F map[int]*SimpleStruct
	G []SimpleStruct
}

func TestComplexStruct(t *testing.T) {
	simpleStruct := SimpleStruct{
		A: 1.,
		B: 2,
		C: "3,",
	}
	simpleStructMap := make(map[int]*SimpleStruct)
	simpleStructMap[1] = &simpleStruct
	simpleStructSlice := []SimpleStruct{simpleStruct}
	complexStruct := ComplexStruct{
		A: 1.,
		B: 2,
		C: "3",
		D: simpleStruct,
		E: &simpleStruct,
		F: simpleStructMap,
		G: simpleStructSlice,
	}
	actualSize := GetMemoryConsumption(complexStruct)
	expectedSize := 216
	if actualSize != expectedSize {
		t.Errorf("Size of complex struct is incorrect (got %v, expected %v)\n", actualSize, expectedSize)
	}
}
