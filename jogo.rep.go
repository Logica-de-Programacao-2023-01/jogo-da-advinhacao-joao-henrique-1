package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numeroOculto := rand.Intn(100) + 1

	var tentativa int
	var numeroTentativas int

	fmt.Println("Bem-vindo ao jogo de adivinhação!")
	fmt.Println("Tente adivinhar o número secreto entre 1 e 100.")

	for {
		fmt.Print("Digite um número: ")
		fmt.Scanln(&tentativa)
		numeroTentativas++

		if tentativa > numeroOculto {
			fmt.Println("O número é menor!")
		} else if tentativa < numeroOculto {
			fmt.Println("O número é maior!")
		} else {
			fmt.Printf("Parabéns! Você acertou o número secreto %d em %d tentativas!\n", numeroOculto, numeroTentativas)
			break
		}
	}
}
