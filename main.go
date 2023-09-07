package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var arabicToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var operandA, operandB int

var operators = map[string]func() int{
	"+": func() int { return operandA + operandB },
	"-": func() int { return operandA - operandB },
	"/": func() int { return operandA / operandB },
	"*": func() int { return operandA * operandB },
}

var inputTokens []string

const ErrInvalidOperation = "Ошибка: Недопустимая математическая операция."

const (
	ErrMixedSystems  = "Ошибка: Смешанные системы счисления не поддерживаются."
	ErrNegativeRoman = "Ошибка: Римские цифры не поддерживают отрицательные числа."
	ErrZeroRoman     = "Ошибка: Римские цифры не представляют ноль."
	ErrOutOfRange    = "Ошибка: Калькулятор работает только с арабскими целыми числами (1-10) или римскими цифрами (I-X)."
)

func performOperation(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romanNumeralsFound := make([]string, 0)
	romanValues := make([]int, 0)

	for op := range operators {
		for _, val := range s {
			if op == string(val) {
				operator += op
				inputTokens = strings.Split(s, operator)
			}
		}
	}

	switch {
	case len(operator) > 1:
		panic(ErrInvalidOperation)
	case len(operator) < 1:
		panic(ErrInvalidOperation)
	}

	for _, elem := range inputTokens {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romanNumeralsFound = append(romanNumeralsFound, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(ErrMixedSystems)
	case 0:
		validInput := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && validInput == true {
			operandA, operandB = numbers[0], numbers[1]
			fmt.Println(val())
		} else {
			panic(ErrOutOfRange)
		}
	case 2:
		for _, elem := range romanNumeralsFound {
			if val, ok := romanNumerals[elem]; ok && val > 0 && val < 11 {
				romanValues = append(romanValues, val)
			} else {
				panic(ErrOutOfRange)
			}
		}
		if val, ok := operators[operator]; ok {
			operandA, operandB = romanValues[0], romanValues[1]
			intToRoman(val())
		}
	}
}

func intToRoman(romanResult int) {
	var romanNumeral string
	if romanResult == 0 {
		panic(ErrZeroRoman)
	} else if romanResult < 0 {
		panic(ErrNegativeRoman)
	}
	for romanResult > 0 {
		for _, elem := range arabicToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range romanNumerals {
					if value == elem {
						romanNumeral += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNumeral)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(input, " ", "")
		performOperation(strings.ToUpper(strings.TrimSpace(s)))
	}
}
