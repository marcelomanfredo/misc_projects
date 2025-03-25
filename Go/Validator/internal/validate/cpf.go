package validate

import (
	"fmt"
	"strconv"

	"Validator/pkg"
)

func ValidateCpf(verbose bool, value string) error {
	if verbose {
		fmt.Println("Validating CPF...")
	}

	// Check if user provided valid lenght
	if len(value) != 11 {
		return fmt.Errorf("CPF string should contain 11 digits")
	}

	// Check first digit
	arrayToValidate := value[0:9]
	check1 := int(value[9] - '0')
	if n, err := utils.CalculateCpfDigit(verbose, arrayToValidate); err != nil {
		return err
	} else if n != check1 {
		fmt.Printf("Invalid CPF: first check digit %d is incorrect.", check1)
		if verbose {
			fmt.Println(" The correct digit is", n)
		} else {
			fmt.Println("") // Adding a 'new line' to output
		}
		return nil
	} else {
		arrayToValidate += strconv.Itoa(n)
	}

	// Check second and last digit
	check2 := int(value[10] - '0')
	if n, err := utils.CalculateCpfDigit(verbose, arrayToValidate); err != nil {
		return err
	} else if n != check2 {
		fmt.Printf("Invalid CPF: second check digit %d is incorrect.", check2)
		if verbose {
			fmt.Println("The correct digit is", n)
		} else {
			fmt.Println("") // Adding a 'new line' to output
		}
		return nil
	}

	fmt.Println("CPF is valid")
	return nil
}
