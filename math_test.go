package main

import "testing"

func TestSoma(t *testing.T) {
	
	total := Soma(1, 2)

	if total != 3 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: 3", total)
	}
}