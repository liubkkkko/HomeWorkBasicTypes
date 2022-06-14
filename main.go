package main

import (
	"fmt"
	"math"
)

func main() {
	const myMoney float64 = 23
	var (
		applePrice   float64 = 5.99
		pearPrice    float64 = 7
		amountApples int     = 9
		amountPears  int     = 8
		solvency     bool
	)
	appleCost := float64(amountApples) * applePrice
	pearCost := float64(amountPears) * pearPrice
	fruitCost := appleCost + pearCost
	applePriceRound := math.Round(applePrice)
	numbersPears := int(myMoney) / int(pearPrice)
	numbersApples := int(myMoney) / int(applePriceRound)
	priceTwoApples := applePrice * 2
	priceTwoPears := pearPrice * 2
	priceFourFruits := priceTwoApples + priceTwoPears
	if priceFourFruits <= myMoney {
		solvency = true
	}
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? --", fruitCost, "грн.")
	fmt.Println("2. Скільки груш ми можемо купити? --", numbersPears, "шт.")
	fmt.Println("3. Скільки яблук ми можемо купити? --", numbersApples, "шт.")
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? --", solvency)

}