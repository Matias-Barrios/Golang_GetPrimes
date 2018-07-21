package main

import (
	"fmt"

	"github.com/Matias-Barrios/getPrimes/modules/utils"
)

func main() {
	for _, primo := range utils.Primes(100) {
		fmt.Println(primo)
	}
}
