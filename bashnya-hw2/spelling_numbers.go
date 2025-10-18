package main

func SpellingHundreds(dig_huns int) string {
	switch dig_huns {
	case 0:
		return ""
	case 1:
		return "сто"
	case 2:
		return "двести"
	case 3:
		return "триста"
	case 4:
		return "четыреста"
	case 5:
		return "пятьсот"
	case 6:
		return "шестьсот"
	case 7:
		return "семьсот"
	case 8:
		return "восемьсот"
	case 9:
		return "девятьсот"
	default:
		return ""
	}
}

func SpellingTens(dig_tens, dig_unis int) string {
	switch dig_tens {
	case 0:
		return ""
	case 1:
		switch dig_unis {
		case 0:
			return "десять"
		case 1:
			return "одиннадцать"
		case 2:
			return "двенадцать"
		case 3:
			return "тринадцать"
		case 4:
			return "четырнадцать"
		case 5:
			return "пятнадцать"
		case 6:
			return "шестнадцать"
		case 7:
			return "семнадцать"
		case 8:
			return "восемнадцать"
		case 9:
			return "девятнадцать"
		default:
			return ""
		}
	case 2:
		return "двадцать"
	case 3:
		return "тридцать"
	case 4:
		return "сорок"
	case 5:
		return "пятьдесят"
	case 6:
		return "шестьдесят"
	case 7:
		return "семьдесят"
	case 8:
		return "восемьдесят"
	case 9:
		return "девяносто"
	default:
		return ""
	}
}

func SpellingUnits(is_thousands bool, dig_huns, dig_tens, dig_ones int) string {
	if dig_tens == 1 {
		return ""
	} else {
		switch dig_ones {
		case 0:
			if !is_thousands && dig_huns == 0 && dig_tens == 0 && dig_ones == 0 {
				return "ноль"
			} else {
				return ""
			}
		case 1:
			if is_thousands {
				return "одна"
			} else {
				return "один"
			}
		case 2:
			if is_thousands {
				return "две"
			} else {
				return "два"
			}
		case 3:
			return "три"
		case 4:
			return "четыре"
		case 5:
			return "пять"
		case 6:
			return "шесть"
		case 7:
			return "семь"
		case 8:
			return "восемь"
		case 9:
			return "девять"
		default:
			return ""
		}
	}
}
