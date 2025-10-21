package main

func SpellingWordThousands(is_more_999 bool, dig_tens, dig_ones int) (end string) {
	if is_more_999 {
		if dig_tens == 1 || dig_ones == 0 || dig_ones >= 5 && dig_ones <= 9 {
			end = " "
		} else if dig_ones == 1 {
			end = "а "
		} else {
			end = "и "
		}
		end = "тысяч" + end
	} else {
		end = ""
	}
	return
}
