package troco

import "testing"

func TestSemTroco(t *testing.T) {
	var preco float32 = 1.0
	var montante float32 = 1.0

	resposta, _ := calcTroco(preco, montante)
	if len(resposta) != 0 {
		t.Error("nao funciona")
	}
}

func TestPrecoMaiorMontante(t *testing.T) {
	var preco float32 = 2.0
	var montante float32 = 1.0

	_, erro := calcTroco(preco, montante)
	if erro {
		t.Error("nao funciona tambem")
	}
}

func TestComTroco(t *testing.T) {
	var preco float32 = 1.0
	var montante float32 = 2.0

	resposta, erro := calcTroco(preco, montante)
	if len(resposta) == 0 && !erro {
		t.Error("nao funciona ainda")
	}
}

func TestComTrocoGrande(t *testing.T) {
	var preco float32 = 100.0
	var montante float32 = 200.0

	resposta, erro := calcTroco(preco, montante)
	if len(resposta) == 0 && !erro {
		t.Error("nao funciona ainda")
	}
}
