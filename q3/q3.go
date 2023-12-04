package q3

//O cozinheiro Remy preparou uma refeição para si mesmo e, enquanto almoçava, decidiu assistir a um vídeo no RataTube. No
//entanto, ele tem um tempo limitado de 𝑡 segundos para o almoço. Ele pediu a sua ajuda para escolher o vídeo ideal.
//
//O RataTube possui um feed de 𝑛 vídeos, cada um representado por uma estrutura de vídeo contendo informações sobre sua
//duração em segundos e o nível de entretenimento. O feed é inicialmente aberto no primeiro vídeo, e Remy pode pular para
//o próximo vídeo em 1 segundo (caso exista). Ele pode pular vídeos quantas vezes desejar, inclusive não pular nenhum.
//
//Sua tarefa é auxiliar Remy a escolher um vídeo que ele possa abrir e assistir dentro do tempo disponível, 𝑡 segundos. Se
//houver vários vídeos que se encaixam nessa condição, ele deseja escolher o vídeo com o maior nível de entretenimento.
//Retorne qualquer vídeo apropriado ou exiba um erro caso não haja um vídeo adequado dentro do tempo disponível.

import (
	"errors"
	"fmt"
)

type Video struct {
	ID            int
	Duration      int
	Entertainment int
}

func ChooseVideo(videos []Video, time int) (Video, error) {
	var selectedVideo Video
	maxEntertainment := -1

	for _, video := range videos {
		if video.Duration <= time {
			if video.Entertainment > maxEntertainment {
				selectedVideo = video
				maxEntertainment = video.Entertainment
			}
		}
	}

	if maxEntertainment == -1 {
		return Video{}, errors.New("nenhum vídeo adequado encontrado")
	}

	return selectedVideo, nil
}

func main() {
	videos := []Video{
		{ID: 1, Duration: 30, Entertainment: 8},
		{ID: 2, Duration: 20, Entertainment: 6},
		{ID: 3, Duration: 40, Entertainment: 9},
		{ID: 4, Duration: 15, Entertainment: 7},
	}

	tempoDisponivel := 25

	videoEscolhido, err := ChooseVideo(videos, tempoDisponivel)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Vídeo escolhido: ID %d, Duração %d segundos, Entretenimento %d\n",
			videoEscolhido.ID, videoEscolhido.Duration, videoEscolhido.Entertainment)
	}
}
