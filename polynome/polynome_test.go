package polynome

import (
	"DMTI_baza/natural"
	"DMTI_baza/rational"
	"DMTI_baza/whole"
	"testing"
)

// Пименов
func TestDerivative2(t *testing.T) {
	var pol, got Polynomial
	var a0, a1, a2, a3 rational.Rational
	var a0n, a1n, a2n, a3n whole.Whole
	var a0d, a1d, a2d, a3d natural.Natural

	var polw Polynomial
	var a0w, a1w, a2w rational.Rational
	var a0nw, a1nw, a2nw whole.Whole
	var a0dw, a1dw, a2dw natural.Natural

	a0n.MakeW(true, []uint8{5})
	a0d.MakeN([]uint8{1})

	a1n.MakeW(false, []uint8{4})
	a1d.MakeN([]uint8{1})

	a2n.MakeW(true, []uint8{3})
	a2d.MakeN([]uint8{1})

	a3n.MakeW(true, []uint8{2})
	a3d.MakeN([]uint8{1})

	a0.MakeR(a0n, a0d)
	a1.MakeR(a1n, a1d)
	a2.MakeR(a2n, a2d)
	a3.MakeR(a3n, a3d)

	pol.MakeP([]rational.Rational{a3, a2, a1, a0})

	a0nw.MakeW(false, []uint8{4})
	a0dw.MakeN([]uint8{1})

	a1nw.MakeW(true, []uint8{6})
	a1dw.MakeN([]uint8{1})

	a2nw.MakeW(true, []uint8{6})
	a2dw.MakeN([]uint8{1})

	a0w.MakeR(a0nw, a0dw)
	a1w.MakeR(a1nw, a1dw)
	a2w.MakeR(a2nw, a2dw)

	polw.MakeP([]rational.Rational{a2w, a1w, a0w})
	got = Derivative(pol)

	if Compare(polw, got) != 0 {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, polw)
	}
}

// Пименов
func TestQuotientOfDivision2(t *testing.T) {
	var pol1, pol2, want, got Polynomial

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(3, 1), rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(3, 1)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(0, 1)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("4 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(2, 1), rational.IntToRational(5, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 2), rational.IntToRational(19, 4)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("5 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 4), rational.IntToRational(1, 2), rational.IntToRational(0, 2)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("6 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(52, 1), rational.IntToRational(96, 1), rational.IntToRational(64, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1), rational.IntToRational(36, 1), rational.IntToRational(104, 1), rational.IntToRational(96, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 4), rational.IntToRational(3, 4)})
	got = QuotientOfDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("7 - получили %+v, ожидалось %+v", got, want)
	}

}

// Пименов
func TestRemainderFromDivision2(t *testing.T) {
	var pol1, pol2, want, got Polynomial

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(0, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(4, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(3, 1), rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(6, 1), rational.IntToRational(4, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("4 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(2, 1), rational.IntToRational(5, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(-79, 4)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("5 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(0, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("6 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(52, 1), rational.IntToRational(96, 1), rational.IntToRational(64, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1), rational.IntToRational(36, 1), rational.IntToRational(104, 1), rational.IntToRational(96, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(-1, 1), rational.IntToRational(-6, 1), rational.IntToRational(-8, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("7 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(4, 1), rational.IntToRational(36, 1), rational.IntToRational(104, 1), rational.IntToRational(96, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(-1, 1), rational.IntToRational(-6, 1), rational.IntToRational(-8, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(0, 1)})
	got = RemainderFromDivision(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("8 - получили %+v, ожидалось %+v", got, want)
	}
}

// Пименов
func TestMultiplicationPol2(t *testing.T) {
	var pol1, pol2, want, got Polynomial

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 4), rational.IntToRational(3, 4)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1), rational.IntToRational(36, 1), rational.IntToRational(104, 1), rational.IntToRational(96, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(53, 1), rational.IntToRational(102, 1), rational.IntToRational(72, 1)})
	got = MultiplicationPol(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}

}

func TestSubstractionP2(t *testing.T) {
	var pol1, pol2, want, got Polynomial

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(52, 1), rational.IntToRational(96, 1), rational.IntToRational(64, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(53, 1), rational.IntToRational(102, 1), rational.IntToRational(72, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(-1, 1), rational.IntToRational(-6, 1), rational.IntToRational(-8, 1)})
	got = SubstractionP(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}
}

// Пименов
func TestGreatestCommonDivisor2(t *testing.T) {
	var pol1, pol2, want, got Polynomial

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(4, 1)})
	got = GreatestCommonDivisor(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(100, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1)})
	got = GreatestCommonDivisor(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(3, 1), rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(2, 1), rational.IntToRational(0, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1)})
	got = GreatestCommonDivisor(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(64, 1), rational.IntToRational(32, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1), rational.IntToRational(2, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(2, 1), rational.IntToRational(1, 1)})
	got = GreatestCommonDivisor(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("4 - получили %+v, ожидалось %+v", got, want)
	}

	pol1.MakeP([]rational.Rational{rational.IntToRational(64, 1), rational.IntToRational(12, 1), rational.IntToRational(4, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(2, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1)})
	got = GreatestCommonDivisor(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("5 - получили %+v, ожидалось %+v", got, want)
	}
	pol1.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(12, 1), rational.IntToRational(52, 1), rational.IntToRational(96, 1), rational.IntToRational(64, 1)})
	pol2.MakeP([]rational.Rational{rational.IntToRational(4, 1), rational.IntToRational(36, 1), rational.IntToRational(104, 1), rational.IntToRational(96, 1)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1), rational.IntToRational(6, 1), rational.IntToRational(8, 1)})
	got = GreatestCommonDivisor(pol1, pol2)
	if Compare(want, got) != 0 {
		t.Fatalf("6 - получили %+v, ожидалось %+v", got, want)
	}
}

// Пименов
func TestGreatestCommonDivisorAndLeastCommonMultipleOfPolynomial2(t *testing.T) {
	var pol, want Polynomial
	var got1, got2 natural.Natural
	pol.MakeP([]rational.Rational{rational.IntToRational(3, 4), rational.IntToRational(1, 4)})
	want.MakeP([]rational.Rational{rational.IntToRational(1, 1)})
	got1, got2 = GreatestCommonDivisorAndLeastCommonMultipleOfPolynomial(pol)
	t.Fatalf("1 - получили %+v %+v, ожидалось %+v", got1, got2, want)
}
