package utils

import (
	"fmt"
	"strconv"

	"marcelomanfredo/misc/Go/logger"
)

func CalculateCpfDigit(value string) (int, error) {
	// Check if array is valid
	if err := validateArray(cpf, value); err != nil {
		return -1, err
	}

	logger.Debugln("\nArray for check: %s", value)
	logger.Debugln("Initializing calculation...")

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

		if i != len(value)-1 {
			logger.Debug("(%d * %d) + ", digit, n+1)
		} else {
			logger.Debug("(%d * %d) = ", digit, n+1)
			logger.Debugln(strconv.Itoa(checkDigit))
		}
	}

	validator := checkDigit * 10 % 11
	logger.Debugln("%d * 10 %% 11 = %d", checkDigit, validator)

	if validator == 10 {
		logger.Debugln("Since the modulus is 10, the validator should be 0")
		validator = 0
	}
	logger.Debugln("Checker digit: %d", validator)

	return validator, nil
}
