package rational

import (
	"DMTI_baza/natural"
	"DMTI_baza/whole"
)

// Rational Структура рационального числа
type Rational struct {
	nominator   whole.Whole
	denominator natural.Natural
}

// MakeR Метод для создания рационального числа
func (r *Rational) MakeR(nom whole.Whole, den natural.Natural) {
	r.nominator = nom
	r.denominator = den
}

func Z_to_q(x whole.Whole) *Rational {
	var one []uint8
	one = append(one, 1)
	p := Rational{nominator: whole.Whole{Negative: false, Num: natural.Natural{Older: 0, Digits: one}}, denominator: natural.Natural{Older: uint32(len(x.Num.Digits)), Digits: x.Num.Digits}}
	return &p
}

func G_to_z(x Rational) *whole.Whole {
	var one []int
	one = append(one, 1)
	p := whole.Whole{Negative: false, Num: x.nominator.Num}
	return &p
}
