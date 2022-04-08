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
	case x.Num.Digits[0] == 0:
		return 0
	case x.Negative == true:
		return 1
	default:
		return 2
	}
}

// MultiplicationByNegativeOne Хвостовский Умножение целого на (-1)
func MultiplicationByNegativeOne(x Whole) Whole {
	x.Negative = !x.Negative
	return x
}

//Комаровский FromNaturalsToWhole преобразование из натурального в целое
func FromNaturalsToWhole(nat natural.Natural) Whole {
	var res Whole
	res.Num = nat
	res.Negative = false
	return res
}

//Комаровский FromWholeToNaturals преобразование из неотрицательного целого в натуральное

func FromWholeToNaturals(wh Whole) natural.Natural {
	var res natural.Natural
	res = wh.Num
	return res

}

//CopyW
//Функция для копирования целого числа
func CopyW(n Whole) Whole {
	var i uint32
	var x Whole
	for i = 0; i <= n.Num.Older; i++ {
		x.Num.Digits = append(x.Num.Digits, n.Num.Digits[i])
	}
	x.Num.Older = n.Num.Older
	x.Negative = n.Negative
	return x
}

//Addition Семёнов
//Сложение двух целых чисел
func Addition(a, b Whole) Whole {
	var fl1, fl2 uint32
	var r, t Whole
	r, t = CopyW(a), CopyW(b)                                                     // коппирую числа
	if natural.Compare(r.Num, t.Num) != 2 && natural.Compare(r.Num, t.Num) != 0 { //Сравниваем числа, если первое небольше второго и они оба не равны, то меняем их местами
		r, t = t, r
	}
	if Positivity(r) == 1 { // запоминаем, если первое отрицательное
		fl1 += 1
	}
	if Positivity(t) == 1 { // запоминаем, если второе отрицательное
		fl2 += 1
	}
	if fl1+fl2 != 1 { // если оба отрицательные или оба положительные складываем цифры
		r.Num = natural.Addition(r.Num, t.Num)
	} else { // иначе вычитаем
		r.Num = natural.Subtraction(r.Num, t.Num)
	}
	return r
}

//Subtraction Семёнов
//Вычитание двух целых чисел
func Subtraction(a, b Whole) Whole {
	var fl1, fl2 uint32
	var r, t Whole
	r, t = CopyW(a), CopyW(b)                                                     // коппирую числа
	if natural.Compare(r.Num, t.Num) != 2 && natural.Compare(r.Num, t.Num) != 0 { //Сравниваем числа, если первое небольше второго и они оба не равны, то меняем их местами
		r, t = t, r
	}
	if Positivity(r) == 1 { // запоминаем, если первое отрицательное
		fl1 += 1
	}
	if Positivity(t) == 1 { // запоминаем, если второе отрицательное
		fl2 += 1
	}
	if fl1+fl2 != 1 { // если оба отрицательные или оба положительные вычитаем цифры
		r.Num = natural.Subtraction(r.Num, t.Num)
	} else { // иначе складываем
		r.Num = natural.Addition(r.Num, t.Num)
	}
	return r
}
