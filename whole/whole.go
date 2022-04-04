package whole

import (
	"DMTI_baza/natural"
)

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
		digits=[0]
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

//Комаровский FromNaturalsToWhole преобразование из натурального в целое
func FromNaturalsToWhole(nat natural.Natural) Whole {
	var res Whole
	res.Num = nat
	res.Negative = false
	return res
}

//FromWholeToNaturals Комаровский преобразование из неотрицательного целого в натуральное
//Если подается неверное условие, выводится true
func FromWholeToNaturals(wh Whole) (natural.Natural, bool) {
	var res natural.Natural
	if wh.Negative == true {
		return nill, true
	}
	else {
		res = wh.Num
		err = 1
		return res, false
	}
}
