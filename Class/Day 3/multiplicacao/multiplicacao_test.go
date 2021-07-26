package multiplicacao

import "testing"

func TestMultiplique(t *testing.T) {
	// chama funcao que multiplica
	obtive := Multiplique(7, 5)
	espero := 35
	// o que espero que aconteça
	if obtive != espero {
		// sinalize que deu ruim, pufavô
		t.Errorf("deu ruim, senhora. espero '%d', mas obtive '%d'", espero, obtive)
	}
	// compara o recebido/esperado
}
