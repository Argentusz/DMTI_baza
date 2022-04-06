package whole

import "DMTI_baza/natural"

// Whole Структура целого числа
type Whole struct {
	Num      natural.Natural
	Negative bool // true (1), если отрицательное; false (0), если положительное
}

// MakeW Метод для создания целого
func (w *Whole) MakeW(Negative bool, digits []uint8) {
	for _, v := range digits {
		if v != 0 { //Проверка на 0 старшей цифры
			break
		} else {
			digits = digits[1:]
		}
	}
	if len(digits) == 0 {
		digits = append(digits, 0)
	}
	w.Num.Digits = digits
	w.Num.Older = uint32(len(digits)) - 1
	w.Negative = Negative
}

// Absolute Функция Тростина Максима
// Возвращает модуль целого числа как натуральное
func Absolute(w Whole) natural.Natural {
	var n natural.Natural
	n.MakeN(w.Num.Digits)
	return n
}

//Positivity Турбина Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)
func Positivity(x Whole) int {
	switch {
	case len(x.Num.Digits) == 0:
		return 0
	case x.Negative == true:
		return 1
	default:
		return 2
	}
}


//Комаровский FromNaturalsToWhole преобразование из натурального в целое
func FromNaturalsToWhole(nat natural.Natural) Whole {
	var res Whole
	res.Num = nat
	res.Negative = false
	return res
}

//Комаровский FromWholeToNaturals преобразование из неотрицательного целого в натуральное(если
//подается отрицательное число то возвращается true)
func FromWholeToNaturals(wh Whole) (natural.Natural, bool) {
	var res natural.Natural
	if wh.Negative == true {
		res.Digits = nil
		return res, true
	} else {
		res = wh.Num
		return res, false
	}
}
// MultiplicationByNegativeOne Хвостовский Умножение целого на (-1)
func MultiplicationByNegativeOne(x Whole) Whole {
	x.Negative = !x.Negative
	return x
}

