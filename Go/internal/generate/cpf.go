package generate

import (
	"fmt"
	"math/rand"
	"strconv"

	"marcelomanfredo/misc/Go/logger"
	"marcelomanfredo/misc/Go/pkg/utils"
)

func GenerateCpf() error {
	logger.Debugln("Generating new CPF...")
	var randValue string
	for range 9 {
		randValue += strconv.Itoa(rand.Intn(9))
	}

	if n1, err := utils.CalculateCpfDigit(randValue); err != nil {
		return err
	} else {
		randValue += strconv.Itoa(n1)
	}

	if n2, err := utils.CalculateCpfDigit(randValue); err != nil {
		return err
	} else {
		randValue += strconv.Itoa(n2)
	}

	logger.Debug("\nRandomly generated CPF: ")
	fmt.Println(randValue)
	return nil
}
