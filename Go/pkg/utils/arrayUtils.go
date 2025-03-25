package utils

import (
	"fmt"
	"unicode"
)

type ArrayType byte

const (
	cpf ArrayType = iota
	mod
)

func validateArray(arrayType ArrayType, array string) error {

	switch arrayType {
	case cpf:
		// Check if every element in the array is a digit
		var digits []byte
		for _, c := range array {
			if !unicode.IsDigit(c) {
				return fmt.Errorf("couldn't convert char %d to int", c)
			}

			var digit byte = byte(c)
			digits = append(digits, digit)
		}

		if allElementsEqual(digits) {
			return fmt.Errorf("every element of the array is equal")
		}

		// Array is valid and the check can go on
		return nil

	case mod:
		return fmt.Errorf("returning error -> mod")
	}
	return fmt.Errorf("array type is not valid")
}

func allElementsEqual(digits []byte) bool {
	var first byte = digits[0]
	for _, n := range digits[1:] {
		if n != first {
			return false
		}
	}
	return true
}
