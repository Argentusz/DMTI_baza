package whole

import "DMTI_baza/natural"

// Whole
// Структура целого числа
type Whole struct {
	Num      natural.Natural
	Negative bool // true (1), если отрицательное; false (0), если положительное
}

// Zero
// Возвращает целый нуль (положительный =) )
func Zero() Whole {
	var zero Whole
	zero.Num = natural.Zero()
	zero.Negative = false
	return zero
}

// ToStringW
// Возвращает целое число как строку
func ToStringW(w Whole) string {
	var s string
	if w.Negative {
		s += "-"
	}
	s += natural.ToStringN(w.Num)
	return s
}

// MakeW
// Метод для создания целого
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

// Absolute Тростина Максима
// Возвращает модуль целого числа как натуральное
func Absolute(w Whole) natural.Natural {
	var n natural.Natural
	n.MakeN(w.Num.Digits)
	return n
}

// Positivity Турбина
// Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)
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

// MultiplicationByNegativeOne Хвостовский
// Умножение целого на (-1)
func MultiplicationByNegativeOne(x Whole) Whole {
	x.Negative = !x.Negative
	return x
}

// FromNaturalsToWhole Комаровский
// Преобразование из натурального в целое
func FromNaturalsToWhole(nat natural.Natural) Whole {
	var res Whole
	res.Num = nat
	res.Negative = false
	return res
}

// FromWholeToNaturals Комаровский
// Преобразование из неотрицательного целого в натуральное
func FromWholeToNaturals(wh Whole) natural.Natural {
	var res natural.Natural
	res = wh.Num
	return res

}

// Multiplication Тростин Максим
// Умножение целых
func Multiplication(x, y Whole) Whole {
	// Если хотя бы один ноль, то возвращаем ноль
	// (чтобы не проходить по всем функциям и не париться по поводу знака)
	if natural.CheckNull(x.Num) || natural.CheckNull(y.Num) {
		return Zero()
	}
	var res Whole
	// Определяем знак результата
	if (x.Negative && y.Negative) || (!x.Negative && !y.Negative) {
		res.Negative = false
	} else {
		res.Negative = true
	}
	// Вычисляем модуль результата
	res.Num = natural.Multiplication(x.Num, y.Num)
	return res
}

//CopyW Семёнов
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
	var r, t, res Whole
	r, t = CopyW(a), CopyW(b)
	if Positivity(r) == 1 { // запоминаем, если первое отрицательное
		fl1 += 1
	}
	if Positivity(t) == 1 { // запоминаем, если второе отрицательное
		fl2 += 1
	}
	if fl1+fl2 != 1 { // если оба отрицательные или оба положительные складываем цифры
		r.Num = natural.Addition(r.Num, t.Num)
		return r
	} else { //иначе рассматриваем их как оба положительных и вычитаем из большего меньшее
		res = CopyW(r) //берём копию первого числа и делаем с ним операции
		res.Num = natural.Subtraction(res.Num, t.Num)
		if natural.Compare(r.Num, t.Num) == 1 && fl1 == 1 { //если изначально первое меньше второго и знак первого минус, то меняем у результата знак на положительный
			res = MultiplicationByNegativeOne(res)
		}
		if natural.Compare(r.Num, t.Num) == 1 && fl2 == 1 { //если изначально первое меньше второго и знак второго минус, меняем на отрицательный
			res = MultiplicationByNegativeOne(res)
		}
	}
	if natural.Compare(a.Num, b.Num) == 0 && fl1 != fl2 { // если числа равны, а их знаки нет, в сумме дают 0, фиксим у этого 0 знак на false
		res.Negative = false
	}
	return res
}

//Subtraction Семёнов
//Вычитание двух целых чисел
func Subtraction(a, b Whole) Whole {
	var fl1, fl2 uint32
	var r, t, res Whole
	r, t = CopyW(a), CopyW(b)
	if Positivity(r) == 1 { // запоминаем, если первое отрицательное
		fl1 += 1
	}
	if Positivity(t) == 1 { // запоминаем, если второе отрицательное
		fl2 += 1
	}
	if fl1 != fl2 { // если оба отрицательные или оба положительные складываем цифры
		r.Num = natural.Addition(r.Num, t.Num)
		return r
	} else { //иначе рассматриваем их как оба положительные и вычитаем из большего меньшее
		res = CopyW(r)
		res.Num = natural.Subtraction(res.Num, t.Num)
		if natural.Compare(r.Num, t.Num) == 1 { //если первое число меньшего второго(не учитываем знак), меняем знак на противоположный
			res = MultiplicationByNegativeOne(res)
		}
	}
	if natural.Compare(a.Num, b.Num) == 0 && fl1 == fl2 { // если числа равны и их знаки тоже,в результате получаем 0, фиксим
		res.Negative = false
	}
	return res
}

//Морозов Никита
//Остаток от деления целого на целое (делитель отличен от нуля)
//На вход: Делимое num1, Делитель num2
func RemainderFromDivision(num1, num2 Whole) Whole {

	//Объявляем результат
	var result Whole

	//Сразу делаем остаток положительным
	result.Negative = false

	//Находим остаток вне зависимости от знаков
	result.Num = natural.RemainderFromDivision(Absolute(num1), Absolute(num2))

	//Если остаток не равен нулю
	if natural.CheckNull(result.Num) == false {
		//Если хотя бы одно число отрицательно, то действуем по формуле
		//result = num2 - (num1 mod num2)
		if Positivity(num1) == 1 || Positivity(num2) == 1 {
			result.Num = natural.Subtraction(Absolute(num2), Absolute(result))
		}
	}

	return result
}

//Морозов Никита
//Частное от деления целого на целое
//На вход: Делмое num1, Делитель num2
func WholeFromDivision(num1, num2 Whole) Whole {
	//Объявляем результат
	var result Whole

	//Находим частное вне зависимости от знаков
	result.Num = natural.IntegerFromDivision(Absolute(num2), Absolute(num1))

	//Если хотя бы один знак отрицателен, прибавялем к частному 1
	if Positivity(num1) == 1 || Positivity(num2) == 1 {
		//Если остаток не равен нулю, то действуем
		if natural.CheckNull(RemainderFromDivision(num1, num2).Num) == false {
			result.Num = natural.Addition1(result.Num)
		}

		//Если отрицательно только одно число, делаем частное отрицательным
		if Positivity(num1) == 1 && Positivity(num2) == 1 {
			result.Negative = false
		} else {
			result.Negative = true
		}
	}

	return result
}
