package generate

import (
	"fmt"
	"math/rand"
	"strconv"

	"Validator/pkg"
)

func GenerateCpf(verbose bool) error {
	if verbose {
		fmt.Println("Generating new CPF...")
	}
	var randValue string
	for range 9 {
		randValue += strconv.Itoa(rand.Intn(9))
	}

	if n1, err := utils.CalculateCpfDigit(verbose, randValue); err != nil {
		return err
	} else {
		if verbose {
			fmt.Println("First check digit:", n1)
		}
		randValue += strconv.Itoa(n1)
	}

	if n2, err := utils.CalculateCpfDigit(verbose, randValue); err != nil {
		return err
	} else {
		if verbose {
			fmt.Println("Second check digit:", n2)
		}
		randValue += strconv.Itoa(n2)
	}

	if verbose {
		fmt.Print("\nRandomly generated CPF: ")
	}
	fmt.Println(randValue)
	return nil
}
