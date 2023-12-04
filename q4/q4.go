package q4

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

import (
	"errors"
	"fmt"
)

func Destino(caminhos [][2]string) (string, error) {
	cidadesPartida := make(map[string]bool)

	for _, caminho := range caminhos {
		cidadesPartida[caminho[0]] = true
	}

	for _, caminho := range caminhos {
		cidadeDestino := caminho[1]
		if _, ok := cidadesPartida[cidadeDestino]; !ok {
			return cidadeDestino, nil
		}
	}

	return "", errors.New("nenhuma cidade de destino encontrada")
}

func main() {
	caminhos := [][2]string{
		{"A", "B"},
		{"B", "C"},
		{"C", "D"},
		{"D", "E"},
	}

	destino, err := Destino(caminhos)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Cidade de destino:", destino)
	}
}
