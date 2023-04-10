package main

import (
	"math/big"
	"strings"
)

func main() {

}

// Переделоть число из 16-й в 10-ю систему счисления
func ConvertHEX2DEC(hex string) string {
	if hex == "" {
		return ""
	}
	hex = strings.ToUpper(hex) // Перевести в верхний регистр

	Big16 := big.NewInt(16) // Постоянный коэффициент

	DecInt := big.NewInt(0)
	// Идём по всему числу
	for _, c := range hex {
		DecInt.Mul(DecInt, Big16) // Домножаем на его разряд

		var increment int64
		switch {
		case '0' <= c && c <= '9':
			increment = int64(c - '0')
		case 'A' <= c && c <= 'F':
			increment = int64(c - 'A' + 10)
		}

		DecInt.Add(DecInt, big.NewInt(increment))
	}

	return DecInt.String()
}
