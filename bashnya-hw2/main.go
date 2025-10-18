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

}
