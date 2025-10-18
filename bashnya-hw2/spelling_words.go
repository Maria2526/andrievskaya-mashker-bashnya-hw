package main

func SpellingWordThousands(is_thousands bool, dig_tens, dig_ones int) (end string) {
	if is_thousands {
		if dig_tens == 1 || dig_ones == 0 || dig_ones >= 5 && dig_ones <= 9 {
			end = ""
		} else if dig_ones == 1 {
			end = "а"
		} else {
			end = "и"
		}
		return "тысяч" + end
	} else {
		return ""
	}
}
