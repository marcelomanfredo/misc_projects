package utils

import (
	"fmt"
)

func CalculateCpfDigit(verbose bool, value string) (int, error) {
	// Check if array is valid
	if err := validateArray(cpf, value); err != nil {
		return -1, err
	}

	if verbose {
		fmt.Println("\nArray for check:", value)
		fmt.Println("Initializing calculation...")
	}

	// CPF check rule:
	// First, we need to multiply the first 9 digits by a sequence of 10..2
	// For the second iteration, we should add 1 to the factor, so the sequence should be 11..2
	// For instance: n[0] * 10 + n[1] * 9 + n[2] * 8 ... n[8] * 2
	//   Or: n[0] * 11 ... n[9] * 2

	var checkDigit int = 0
	var n byte
	switch len(value) {
	case 9:
		// First iteration
		n = 10
	case 10:
		// Second iteration
		n = 11
	default:
		return -1, fmt.Errorf("something went wrong while calculating check digits")
	}

	for i, c := range value {
		digit := byte(c) - '0'
		checkDigit += int(digit * n)
		n -= 1
		if verbose {
			if i != len(value)-1 {
				fmt.Printf("(%d * %d) + ", digit, n+1)
			} else {
				fmt.Printf("(%d * %d) = ", digit, n+1)
				fmt.Println(checkDigit)
			}
		}
	}

	validator := checkDigit * 10 % 11
	if verbose {
		fmt.Printf("%d * 10 %% 11 = %d\n", checkDigit, validator)
	}
	if validator == 10 {
		if verbose {
			fmt.Printf("Since the modulus is 10, the validator should be 0\n")
		}
		validator = 0
	}
	return validator, nil
}
