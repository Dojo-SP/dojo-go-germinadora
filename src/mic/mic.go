package main

/*
  Retorna verdadeiro ou falso para disponibilidade
*/
func PossoMijar(ocupados []bool) bool {

	/*for i := 0; i < len(ocupados); i++ {
		if i != 0 {
			disp[i] = !ocupados[i]
			if len(ocupados) > 1 {
				disp[i+1] = !ocupados[i]
			}
		}
		disp[i] = !ocupados[i]
	}*/

	for _, mic := range ocupados {
		if mic {
			return false
		}
	}
	return true
}

func main() {
}
