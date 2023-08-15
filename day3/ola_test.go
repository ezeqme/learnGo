package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em francês", func(t *testing.T) {
		resultado := Ola("Zac", "frances")
		esperado := "Bonjour, Zac"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em esperanto", func(t *testing.T) {
		resultado := Ola("mondo", "esperanto")
		esperado := "Saluton, mondo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
