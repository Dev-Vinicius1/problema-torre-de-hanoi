package main

import "fmt"

func moveHaste(haste int, origem, destino, aux string) {
	if haste == 1 {
		fmt.Printf("Movendo disco de %s para %s\n", origem, destino)

		return
	}

	moveHaste(haste-1, origem, aux, destino)
	fmt.Printf("Movendo disco de %s para %s\n", origem, destino)
	moveHaste(haste-1, aux, destino, origem)
}

func main() {
	hastes := 3 // CUIDADO: O(2^n).
	moveHaste(hastes, "[TorreA]", "[TorreB]", "[TorreC]")
}
