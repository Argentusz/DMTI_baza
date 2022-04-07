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
	zero.Denominator = natural.Natural{Digits: []uint8{1}, Older: 1}
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
	return x.Denominator.Older == 0 && x.Denominator.Digits[0] == 1

}
