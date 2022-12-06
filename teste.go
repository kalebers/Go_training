package main

//todo programa em GO tbm roda  na funcao main

// := declara e inicializa uma variável. forma curta de dizer var x = something

//caso queira fazer mais imports, posso apenas colocar dentro dos parenteses
import (
	"fmt"
	"time"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello")
	fmt.Println(time.Now())

	// teste com funcoes
	fmt.Println("Funcoes")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("Teste com constantes numericas")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
