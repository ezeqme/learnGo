package main

import "testing"

// Para escrever um teste em Go é como escrever uma função
// precisa estar num arquivo parecido com xxx_test.go
// a função precisa começar com a palavra Test
// a função de teste recebe um único argumento que é t *testing.T

func TestHello(t *testing.T) {
	resultado := Hello()
	esperado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("resultado %q, esperado %q", resultado, esperado)
	}
}
