package main

import (
	"fmt"
	"strings"
)

const (
	euro float64 = 94.04
	usd  float64 = 83.92
)

func rubToUsd(rub float64) float64  { return rub / usd }
func rubToEuro(rub float64) float64 { return rub / euro }
func usdToEuro(d float64) float64   { return rubToEuro(d * usd) }

func main() {
	for {
		var rub float64
		fmt.Print("введите сумму в рублях: ")
		fmt.Scan(&rub)

		fmt.Printf("\n%.2f USD\n", rubToUsd(rub))
		fmt.Printf("%.2f EUR\n", rubToEuro(rub))
		fmt.Printf("%.2f EUR из 100 USD\n", usdToEuro(100))

		var answer string
		fmt.Print("\nсконвертировать ещё? (yes/no): ")
		fmt.Scan(&answer)

		if strings.ToLower(answer) == "no" || strings.ToLower(answer) == "n" {
			fmt.Println("end")
			break
		}
	}
}
