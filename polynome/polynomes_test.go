package polynome

import (
	"DMTI_baza/natural"
	"DMTI_baza/rational"
	"DMTI_baza/whole"
	"testing"
)

func TestAdditionP(t *testing.T) {
	var p1, p2, pWant Polynomial

	p1.Coeff = append(p1.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{1}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	p1.Coeff = append(p1.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{1, 7}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{2, 6}, Older: 1}})
	p2.Coeff = append(p2.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{4}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	p2.Coeff = append(p2.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{3, 4}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1, 3}, Older: 1}})
	p2.Coeff = append(p2.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{0}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	p1.Older = 1
	p2.Older = 2
	pWant.Coeff = append(pWant.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{4}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	pWant.Coeff = append(pWant.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{4, 7}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1, 3}, Older: 1}})
	pWant.Coeff = append(pWant.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{1, 7}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{2, 6}, Older: 1}})
	pWant.Older = 2
	got := AdditionP(p1, p2)
	want := pWant
	for i, v := range got.Coeff {
		if v.Nominator.Num.Digits[i] != want.Coeff[i].Nominator.Num.Digits[i] {
			t.Fatalf("%+v got,%+v want", v, want.Coeff[i])
		}
	}
}

func TestSubstractionP(t *testing.T) {
	var p1, p2, pWant Polynomial

	p1.Coeff = append(p1.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{1}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	p1.Coeff = append(p1.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{1, 7}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{2, 6}, Older: 1}})
	p2.Coeff = append(p2.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{4}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	p2.Coeff = append(p2.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{3, 4}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1, 3}, Older: 1}})
	p2.Coeff = append(p2.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{0}, Older: 0}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	p1.Older = 1
	p2.Older = 2
	pWant.Coeff = append(pWant.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{4}, Older: 0}, Negative: true},
		Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}})
	pWant.Coeff = append(pWant.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{2, 1}, Older: 1}, Negative: true},
		Denominator: natural.Natural{Digits: []uint8{1, 3}, Older: 1}})
	pWant.Coeff = append(pWant.Coeff, rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{1, 7}, Older: 1}, Negative: false},
		Denominator: natural.Natural{Digits: []uint8{2, 6}, Older: 1}})
	got := SubstractionP(p1, p2)
	want := pWant
	for i, v := range got.Coeff {
		if v.Nominator.Num.Digits[i] != want.Coeff[i].Nominator.Num.Digits[i] {
			t.Fatalf("%+v got,%+v want", v, want.Coeff[i])
		}
	}
}
