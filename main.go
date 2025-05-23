package main

import (
	"fmt"
	"strconv"
	"strings"
)

const USDtoEUR = 0.89
const USDtoRUB = 80.85
const EURtoRUB = USDtoRUB / USDtoEUR

func main() {
	currencyNames := []string{"USD", "EUR", "RUB"}
	initCurrencyName := geNameCurrencyFromUser(true, currencyNames)
	targetCurrencyNames := getSliceTargetCurrencyNames(currencyNames, initCurrencyName)
	amountOfCurrency := getAmountOfCurrency()
	targetCurrencyName := geNameCurrencyFromUser(false, targetCurrencyNames)
	sum := calculationCurrency(amountOfCurrency, initCurrencyName, targetCurrencyName)
	fmt.Printf("%.2f %s стоит %.2f %s\n", amountOfCurrency, initCurrencyName, sum, targetCurrencyName)
}

func getSliceTargetCurrencyNames(slice []string, initCurrencyName string) []string {
	formatName := strings.ToUpper(initCurrencyName)
	s := make([]string, 0, 2)
	for _, v := range slice {
		if formatName != v {
			s = append(s, v)
		}
	}
	return s
}

func geNameCurrencyFromUser(isInitialCurrency bool, slice []string) string {
	currencyNames := strings.Join(slice, " ")
	var sourceCurrency string
	var currencyAttribute string
	var s1 string

	if isInitialCurrency {
		currencyAttribute = "исходную"
		s1 = "исходной"
	} else {
		currencyAttribute = "целевую"
		s1 = "целевой"
	}

	for {
		fmt.Printf("Введите %s валюту (%s): ", currencyAttribute, currencyNames)
		fmt.Scan(&sourceCurrency)
		formatStr := strings.ToUpper(sourceCurrency)

		if contains(slice, formatStr) {
			return formatStr
		} else {
			fmt.Printf("Введен неверный формат %s валюты. Повторите ввод.\n", s1)
			continue
		}
	}
}

func getAmountOfCurrency() float64 {
	for {
		var input string
		fmt.Print("Введите количество валюты: ")
		fmt.Scan(&input)
		num, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Ошибка: введите число")
			continue
		} else if num < 0 {
			fmt.Println("Ошибка: введите число больше 0")
			continue
		} else {
			return num
		}
	}
}

func calculationCurrency(num float64, init string, target string) float64 {
	switch {
	case init == "USD" && target == "EUR":
		return USDtoEUR * num
	case init == "USD" && target == "RUB":
		return USDtoRUB * num
	case init == "EUR" && target == "USD":
		return num / USDtoEUR
	case init == "EUR" && target == "RUB":
		return EURtoRUB * num
	case init == "RUB" && target == "USD":
		return num / USDtoRUB
	case init == "RUB" && target == "EUR":
		return num / EURtoRUB
	}
	return 0.0
}

func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
