package q2

//O torneio de programação do CEUB ocorrerá em breve. Neste ano, equipes de quatro pessoas estão autorizadas a participar.
//
//No UniCEUB, temos um grupo de participantes que inclui programadores e matemáticos. Gostaríamos de saber o número máximo
//de equipes que podem ser formadas, considerando as seguintes regras:
//
//- Cada equipe deve ser composta por exatamente quatro estudantes.
//- Equipes compostas apenas por quatro matemáticos ou apenas por quatro programadores não têm um bom desempenho,
//  portanto, decidiu-se não formar tais equipes.
//- Assim, cada equipe deve ter pelo menos um programador e pelo menos um matemático.
//
//Escreva um programa que receba como entrada uma lista de participantes e retorne o número máximo de equipes que podem
//ser formadas, respeitando as regras mencionadas.
//
//Cada pessoa só pode fazer parte de uma equipe.

import "fmt"

type Participant struct {
	Name string
	Role string
}

func CalculateTeams(participants []Participant) int {
	numProgrammers := 0
	numMathematicians := 0

	teamsFormed := 0

	for _, participant := range participants {
		switch participant.Role {
		case "Programador":
			numProgrammers++
		case "Matemático":
			numMathematicians++
		}
	}

	for numProgrammers >= 1 && numMathematicians >= 1 {
		numProgrammers -= 1
		numMathematicians -= 1
		teamsFormed++
	}

	return teamsFormed
}

func main() {
	participantes := []Participant{
		{Name: "Alice", Role: "Programador"},
		{Name: "Bob", Role: "Matemático"},
		{Name: "Charlie", Role: "Matemático"},
		{Name: "David", Role: "Programador"},
		{Name: "Eva", Role: "Programador"},
	}

	numEquipes := CalculateTeams(participantes)

	fmt.Printf("Número máximo de equipes formadas: %d\n", numEquipes)
}
