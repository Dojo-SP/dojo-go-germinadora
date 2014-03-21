package troco

func calcTroco(preco float32, montante float32) ([]float32, bool) {
	var cedulasDisponiveis []float32 = []float32{100.0, 50.0, 10.0, 5.0, 1.0, 0.5, 0.1, 0.05, 0.01}

	var cedulasTroco []float32

	if preco > montante {
		return cedulasTroco, false
	} else if preco == montante {
		return cedulasTroco, true
	} else {
		valortroco := montante - preco

		for _, cedula := range cedulasDisponiveis {
			if valortroco == 0 {
				break
			}

			if valortroco-cedula >= 0 {
				cedulasTroco = append(cedulasTroco, cedula)
				valortroco -= cedula
			}
		}
		return cedulasTroco, true
	}
}
