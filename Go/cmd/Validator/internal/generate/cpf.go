package generate

import (
	"fmt"
	"math/rand"
	"strings"

	"marcelomanfredo/misc/Go/logger"
	"marcelomanfredo/misc/Go/pkg/utils"
)

func GenerateCpf() error {
	logger.Debugln("Generating new CPF...")
	var s strings.Builder
	for range 9 {
		s.WriteByte('0' + byte(rand.Intn(10)))
	}

	if n1, err := utils.CalculateCpfDigit(s.String()); err != nil {
		return err
	} else {
		s.WriteByte('0' + byte(n1))
	}

	if n2, err := utils.CalculateCpfDigit(s.String()); err != nil {
		return err
	} else {
		s.WriteByte('0' + byte(n2))
	}

	logger.Debug("\nRandomly generated CPF: ")
	fmt.Println(s.String())
	return nil
}
