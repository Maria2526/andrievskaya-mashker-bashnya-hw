package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var n int
	for {
		fmt.Print("Введите целое число: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		num, err := strconv.Atoi(input)
		if err == nil {
			if num >= 12307 {
				fmt.Println(input, "не подходит, оно не меньше 12307")
			} else if num*-1 > 12307 {
				fmt.Println(input, "не подходит, введите число меньшее по модулю") // иначе там любое отрицательное число можно ввести, вывод бесконечный
			} else {
				n = num
				break
			}
		} else {
			fmt.Println("Ошибка:", input, "не является целым числом")
		}
	}

	for n < 12307 {
		if n < 0 {
			n *= -1
		} else if n%7 == 0 {
			n *= 39
		} else if n%9 == 0 {
			n = n*13 + 1
		} else {
			n = (n + 2) * 3
		}

		if n%13 == 0 && n%9 == 0 {
			fmt.Print("service error")
			break
		} else {
			n++
		}
	}

	digit_units := n % 10
	digit_tens := n / 10 % 10
	digit_hndrds := n / 100 % 10

	digit_units_th := n / 1000 % 10
	digit_tens_th := n / 10000 % 10
	digit_hndrds_th := n / 100000

	var spelling string
	spelling += SpellingHundreds(digit_hndrds_th) + " "
	spelling += SpellingTens(digit_tens_th, digit_units_th) + " "
	spelling += SpellingUnits(true, digit_hndrds_th, digit_tens_th, digit_units_th) + " "
	spelling += SpellingWordThousands(true, digit_tens_th, digit_units_th) + " "
	spelling += SpellingHundreds(digit_hndrds) + " "
	spelling += SpellingTens(digit_tens, digit_units) + " "
	spelling += SpellingUnits(false, digit_hndrds, digit_tens, digit_units)

	fmt.Printf("Итоговое число: %d\nЕго запись: %s", n, spelling)

}
