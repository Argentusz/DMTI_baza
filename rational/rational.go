package rational

import (
	"DMTI_baza/natural"
	"DMTI_baza/whole"
	"fmt"
)

// Rational
// Структура рационального числа
type Rational struct {
	Nominator   whole.Whole
	Denominator natural.Natural
}

// Zero
// Возвращает рациональный нуль
func Zero() Rational {
	var zero Rational
	zero.Nominator = whole.Zero()
	zero.Denominator = natural.Natural{Digits: []uint8{1}, Older: 0}
	return zero
}

// ToStringR
// Возвращает Rational как строку
func ToStringR(r Rational) string {
	var s string
	if r.Nominator.Negative {
		s += "-"
	}
	s += whole.ToStringW(r.Nominator)
	s += fmt.Sprint("/")
	s += natural.ToStringN(r.Denominator)
	return s
}

// MakeR
// Метод для создания рационального числа
func (r *Rational) MakeR(nom whole.Whole, den natural.Natural) {
	r.Nominator = nom
	r.Denominator = den
}

// WholeToFractional Грунской Натальи
// Функция перевода целого в дробное
func WholeToFractional(x whole.Whole) Rational {
	var one []uint8
	var p Rational
	one = append(one, 1) //делаем знаменатель 1

	p.Nominator = whole.Whole{Num: natural.Natural{Digits: x.Num.Digits, //в числитель переносим массив цифр и знак
		Older: uint32(len(x.Num.Digits))},
		Negative: x.Negative}

	p.Denominator = natural.Natural{Digits: one, //в знаменатель переносим единичку
		Older: 0}
	return p
}

// FractionalToWhole Грунской Натальи
// Функция перевода дробного числа в целое
func FractionalToWhole(x Rational) whole.Whole {

	p := whole.Whole{Num: x.Nominator.Num, //переносим только числитель т.к. знаменатель 1
		Negative: x.Nominator.Negative} //переносим знак числителя

	return p
}

// CheckingForWhole Комаровский
// Проверка рационального числа на целое
func CheckingForWhole(x Rational) bool {
	var res Rational
	res = CopyR(x)
	res = SimplifyingFractions(res)
	return res.Denominator.Older == 0 && res.Denominator.Digits[0] == 1

}

//CopyW Семёнов
//Функция для копирования целого числа
func CopyR(n Rational) Rational {
	var i uint32
	var x Rational
	for i = 0; i <= n.Nominator.Num.Older; i++ {
		x.Nominator.Num.Digits = append(x.Nominator.Num.Digits, n.Nominator.Num.Digits[i])
	}
	for i = 0; i <= n.Denominator.Older; i++ {
		x.Denominator.Digits = append(x.Denominator.Digits, n.Denominator.Digits[i])
	}
	x.Nominator.Num.Older = n.Nominator.Num.Older
	x.Denominator.Older = n.Denominator.Older
	x.Nominator.Negative = n.Nominator.Negative
	return x
}

//SimplifyingFractions Семёнов
//Функция скоращения дробей
func SimplifyingFractions(a Rational) Rational {
	var NOD natural.Natural
	var Copy Rational
	DigEd := []uint8{1}
	ed := natural.Natural{DigEd, 0}
	Copy = CopyR(a) //делаю копию

	if Copy.Nominator.Num.Digits[0] == 0 || Copy.Denominator.Digits[0] == 0 { // если в знаменателе//числителе 0, то возвращаем функцию, которую мне сказали возвращать, простите
		return Zero()
	}

	NOD = natural.GreatestCommonDivisor(Copy.Nominator.Num, Copy.Denominator) // нахожу наибольший общий делитель

	if natural.Compare(ed, NOD) == 0 { // если NOD == 1, то возврашаю дробь, не деля её
		return Copy
	} else { // если нет, то делю
		Copy.Denominator = natural.IntegerFromDivision(Copy.Denominator, NOD)     // делю на НОД числитель
		Copy.Nominator.Num = natural.IntegerFromDivision(Copy.Nominator.Num, NOD) // делю на НОД знаменатель
	}

	return Copy
}

// Addition Комаровский
//Сложение дробей
func Addition(x1, x2 Rational) Rational {
	var res, a, b Rational
	var nod, d1, d2 natural.Natural // d1,d2 множители необходимы для приведению к общему знаменателю
	a, b = CopyR(x1), CopyR(x2)
	if natural.Compare(a.Denominator, b.Denominator) == 0 { //если знаменатели равны просто складываем числители
		res.Nominator = whole.Addition(a.Nominator, b.Nominator)
		res.Denominator = a.Denominator
		res = SimplifyingFractions(res)
		return res
	}
	nod = natural.LeastCommonMultiple(a.Denominator, b.Denominator) // Находим Нод знаменателей
	d1 = natural.IntegerFromDivision(b.Denominator, nod)            // ищем недостающие множители для приведению к общему знаменателю
	d2 = natural.IntegerFromDivision(a.Denominator, nod)
	res.Denominator = natural.Multiplication(a.Denominator, d2)   // находим новый знаменатель
	a.Nominator.Num = natural.Multiplication(a.Nominator.Num, d2) // Находим новые числители
	b.Nominator.Num = natural.Multiplication(b.Nominator.Num, d1)
	res.Nominator = whole.Addition(a.Nominator, b.Nominator) // складываем числители
	res = SimplifyingFractions(res)
	return res
}

// Subtraction Комаровский
// Вычитание дробей
func Subtraction(x1, x2 Rational) Rational {
	var res, a, b Rational
	var nod, d1, d2 natural.Natural
	a, b = CopyR(x1), CopyR(x2)
	if natural.Compare(a.Denominator, b.Denominator) == 0 { //если знаменатели равны просто вычитаем числители
		res.Nominator = whole.Subtraction(a.Nominator, b.Nominator)
		res.Denominator = a.Denominator
		res = SimplifyingFractions(res)
		return res
	}
	nod = natural.LeastCommonMultiple(a.Denominator, b.Denominator) // Находим Нод знаменателей
	d1 = natural.IntegerFromDivision(b.Denominator, nod)            // ищем недостающие множители для приведению к общему знаменателю
	d2 = natural.IntegerFromDivision(a.Denominator, nod)
	res.Denominator = natural.Multiplication(a.Denominator, d2)   // находим новый знаменатель
	a.Nominator.Num = natural.Multiplication(a.Nominator.Num, d2) // Находим новые числители
	b.Nominator.Num = natural.Multiplication(b.Nominator.Num, d1)
	res.Nominator = whole.Subtraction(a.Nominator, b.Nominator) // вычитаем числители
	res = SimplifyingFractions(res)
	return res

}
