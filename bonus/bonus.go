package bonus

//Receba uma lista de camisas, cada uma contendo o preço e o tamanho. O tamanho da camisa é representado por uma string,
//que pode ser "M" ou uma combinação de caracteres "X" seguida por "S" ou "L".
//
//Por exemplo, as strings "M", "XXL", "S" e "XXXXXXXS" podem representar tamanhos de algumas camisas. Já as strings "
//XM", "LL" e "SX" não representam tamanhos válidos.
//
//O objetivo é calcular a média entre o preço da maior camisa e o preço da menor camisa da lista.
//
//A comparação entre os tamanhos das camisas deve seguir as seguintes regras:
//
//- Qualquer tamanho pequeno (independentemente da quantidade de caracteres "X") é menor que o tamanho médio e qualquer
//  tamanho grande.
//- Qualquer tamanho grande (independentemente da quantidade de caracteres "X") é maior que o tamanho médio e qualquer
//  tamanho pequeno.
//- Quanto mais caracteres "X" antes de "S", menor o tamanho.
//- Quanto mais caracteres "X" antes de "L", maior o tamanho.
//
//Por exemplo:
//
//1. "XXXS" < "XS"
//2. "XXXL" > "XL"
//3. "XL" > "M"
//4. "XXL" = "XXL"
//5. "XXXXXS" < "M"
//6. "XL" > "XXXS"
//
//Dessa forma, ao receber a lista de camisas com seus respectivos preços e tamanhos, você deve calcular a média de preços
//da maior e da menor camisa.
//
//Caso não seja possível calcular a média, retorne um erro.

import (
	"errors"
	"fmt"
	"math"
)

type Shirt struct {
	Size  string
	Price float64
}

func CalculateAveragePrice(shirts []Shirt) (max float64, min float64, err error) {
	if len(shirts) == 0 {
		return 0, 0, errors.New("a lista de camisas está vazia")
	}

	max = shirts[0].Price
	min = shirts[0].Price

	for _, shirt := range shirts {
		if shirt.Price > max {
			max = shirt.Price
		}

		if shirt.Price < min {
			min = shirt.Price
		}
	}

	return max, min, nil
}

func main() {
	camisas := []Shirt{
		{Size: "M", Price: 25.0},
		{Size: "XXL", Price: 30.0},
		{Size: "S", Price: 20.0},
		{Size: "XXXXXXXS", Price: 15.0},
	}
	
	maxPrice, minPrice, err := CalculateAveragePrice(camisas)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		average := (maxPrice + minPrice) / 2
		fmt.Printf("Média de preços: %.2f\n", average)
	}
}
