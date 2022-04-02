package rational

import (
	"DMTI_baza/natural"
	"DMTI_baza/whole"
)

// Rational Структура рационального числа
type Rational struct {
	Nominator   whole.Whole
	Denominator natural.Natural
}

// MakeR Метод для создания рационального числа
func (r *Rational) MakeR(nom whole.Whole, den natural.Natural) {
	r.Nominator = nom
	r.Denominator = den
}

//Функия Грунской Натальи
//Функция перевода целого в дробное

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

//Функия Грунской Натальи
//Функция перевода дробного числа в целое

func FractionalToWhole(x Rational) whole.Whole {

	p := whole.Whole{Num: x.Nominator.Num, //переносим только числитель т.к. знаменатель 1
		Negative: x.Nominator.Negative} //переносим знак числителя

	return p
}
