package q1

//Na loja de animais à venda, existem alguns tipos de ração disponíveis para compra, sendo eles:
//
//* Ração para cachorro
//* Ração para gato
//* Ração universal
//
//O estoque dessas rações está representado em um mapa, onde a chave é o nome da ração e o valor correspondente é a
//quantidade disponível em estoque.
//
//Polycarp possui 𝑥 cães e 𝑦 gatos. Gostaríamos de determinar se é possível para ele comprar comida suficiente para todos
//os seus animais na loja. Cada um dos seus cães e gatos deve receber um pacote de ração adequado para sua espécie.

import "fmt"

func CanBuyFood(stock map[string]int, dogs, cats int) bool {
	if stock["Ração para cachorro"] < dogs {
		return false
	}

	if stock["Ração para gato"] < cats {
		return false
	}

	return true
}

func main() {
	estoque := map[string]int{
		"Ração para cachorro": 10,
		"Ração para gato":     15,
		"Ração universal":     20,
	}

	numCachorros := 5
	numGatos := 8

	podeComprar := CanBuyFood(estoque, numCachorros, numGatos)

	if podeComprar {
		fmt.Println("Polycarp pode comprar comida suficiente para seus animais.")
	} else {
		fmt.Println("Polycarp não pode comprar comida suficiente para todos os seus animais.")
	}
}
