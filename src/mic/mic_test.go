package main

import (
	"testing"
)

func TestUmMictorioNenhumMijao(t *testing.T) {
	ocupados := []bool{false}
	var vagas = PossoMijar(ocupados)

	if vagas != true {
		t.Error("Deveria retornar uma vaga segura!")
	}
}

func TestUmMictorioUmMijao(t *testing.T) {
	ocupados := []bool{true}
	var vagas = PossoMijar(ocupados)

	if vagas != false {
		t.Error("Teste com um mictorio e um mijao falhou!")
	}
}

func TestDoisMictoriosNenhumMijao(t *testing.T){
	ocupados := []bool{false,false}
	var vagas = PossoMijar(ocupados)

	if !vagas {
		t.Error("Falhou TestDoisMictoriosNenhumMijao")
	}
}

func TestDoisMictoriosComMijao(t *testing.T) {

	var vagasPrimeiroMic = PossoMijar([]bool{false,true})
	var vagasUltimoMic = PossoMijar([]bool{true,false})
	var vagasDoisMics = PossoMijar([]bool{true,true})

	if vagasPrimeiroMic {
		t.Error("Falhou TestDoisMictoriosComMijao <")
	}

	if  vagasUltimoMic {
		t.Error("Falhou TestDoisMictoriosComMijao >")	
	}

	if  vagasDoisMics {
		t.Error("Falhou TestDoisMictoriosComMijao ..")	
	}
}

func TestTresMictoriosSemMijao(t *testing.T) {
	ocupados := []bool{false,false, false}
	var vagas = PossoMijar(ocupados)

	if !vagas {
		t.Error("Falhou TestTresMictoriosSemMijao")
	}
}

func TestTresMictoriosComUmMijao(t *testing.T) {
	var temVagas = PossoMijar([]bool{false,true, false})
	var vagasUltimoMic = PossoMijar([]bool{true,false, false})
	var vagasDoisMics = PossoMijar([]bool{false, false, true})

	if !temVagas {
		t.Error("Falhou TestDoisMictoriosComMijao <")
	}

	if !vagasUltimoMic {
		t.Error("Falhou TestDoisMictoriosComMijao >")	
	}

	if  vagasDoisMics {
		t.Error("Falhou TestDoisMictoriosComMijao ..")	
	}
}

/*
func TestDoisMictoriosAlgumMijao(t *testing.T) {
	ocupados := [2][2]bool{}
	append(ocupados, [2]bool{true, false})
	append(ocupados, [2]bool{true, false})
	var vagas bool 
	for _, value := range(ocupados) {
		vagas = PossoMijar(value)
		if !vagas {
			t.Error("Falhou TestDoisMictoriosAlgumMijao")
		} 
	}

}
*/