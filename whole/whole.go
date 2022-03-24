package whole

import "DMTI_baza/natural"

// Whole Структура целого числа
type Whole struct {
	Digits   []uint8
	Older    uint32
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
	w.Digits = digits
	w.Older = uint32(len(digits))
	w.Negative = Negative
}

// Abs Функция Тростина Максима
// возвращает модуль целого числа как натуральное
func Abs(w Whole) natural.Natural {
	var n natural.Natural
	n.MakeN(w.Digits)
	return n
}
