package validate

import (
	"fmt"
	"strconv"

	"marcelomanfredo/misc/Go/logger"
	"marcelomanfredo/misc/Go/pkg/utils"
)

func ValidateCpf(value string) error {
	logger.Debugln("Validating CPF \"%s\"...", value)

	// Check if user provided valid lenght
	if len(value) != 11 {
		return fmt.Errorf("CPF string should contain 11 digits")
	}

	// Check first digit
	arrayToValidate := value[0:9]
	check1 := int(value[9] - '0')
	if n, err := utils.CalculateCpfDigit(arrayToValidate); err != nil {
		return err
	} else if n != check1 {
		fmt.Printf("\nInvalid CPF: first check digit %d is incorrect.", check1)
		logger.Debug(" The correct digit is %d", n)
		fmt.Println("") // Adding a 'new line' to output
		return nil
	} else {
		arrayToValidate += strconv.Itoa(n)
	}

	// Check second and last digit
	check2 := int(value[10] - '0')
	if n, err := utils.CalculateCpfDigit(arrayToValidate); err != nil {
		return err
	} else if n != check2 {
		fmt.Printf("\nInvalid CPF: second check digit %d is incorrect.", check2)
		logger.Debug(" The correct digit is %d", n)
		fmt.Println("") // Adding a 'new line' to output
		return nil
	}

	fmt.Println("CPF is valid")
	return nil
}
