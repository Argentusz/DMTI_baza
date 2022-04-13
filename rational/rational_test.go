package rational

import (
	"DMTI_baza/natural"
	"DMTI_baza/whole"
	"testing"
)

// Пименов
func TestAddition(t *testing.T) {
	var n1, n2, got, want Rational
	var nom1, nom2, wantn whole.Whole
	var den1, den2, wantd natural.Natural
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 3})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{1, 3})
	den2.MakeN([]uint8{1, 3})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{2})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{2, 6})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{1, 3})
	den2.MakeN([]uint8{1, 3})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{3})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{1, 1})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{2, 7, 6})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 3})
	den1.MakeN([]uint8{7})

	nom2.MakeW(false, []uint8{1, 9})
	den2.MakeN([]uint8{1, 1})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{2, 7, 6})
	wantd.MakeN([]uint8{7, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{1, 9, 8})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("4 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(true, []uint8{1, 1})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(true, []uint8{1, 0})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("5 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 3})
	den1.MakeN([]uint8{7})

	nom2.MakeW(true, []uint8{1, 9})
	den2.MakeN([]uint8{1, 1})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{1, 0})
	wantd.MakeN([]uint8{7, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("6 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(true, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(false, []uint8{6, 8})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("7 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(true, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(true, []uint8{1, 1})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Addition(n1, n2)
	wantn.MakeW(true, []uint8{2, 7, 6})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("8 - получили %+v, ожидалось %+v", got, want)
	}
}

// Пименов
func TestSubtraction(t *testing.T) {
	var n1, n2, got, want Rational
	var nom1, nom2, wantn whole.Whole
	var den1, den2, wantd natural.Natural
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 3})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{1, 3})
	den2.MakeN([]uint8{1, 3})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{0})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{2, 6})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{1, 3})
	den2.MakeN([]uint8{1, 3})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{1})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{1, 1})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(true, []uint8{1, 0})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 3})
	den1.MakeN([]uint8{7})

	nom2.MakeW(false, []uint8{1, 9})
	den2.MakeN([]uint8{1, 1})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{1, 0})
	wantd.MakeN([]uint8{7, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{6, 8})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("4 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(true, []uint8{1, 1})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{2, 7, 6})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("5 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 3})
	den1.MakeN([]uint8{7})

	nom2.MakeW(true, []uint8{1, 9})
	den2.MakeN([]uint8{1, 1})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{2, 7, 6})
	wantd.MakeN([]uint8{7, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("6 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(true, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{1, 9, 8})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("7 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(true, []uint8{7})
	den1.MakeN([]uint8{1, 3})

	nom2.MakeW(true, []uint8{1, 1})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(false, []uint8{1, 0})
	wantd.MakeN([]uint8{2, 4, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("8 - получили %+v, ожидалось %+v", got, want)
	}
	////////////////////////////////////////
	nom1.MakeW(false, []uint8{6, 4})
	den1.MakeN([]uint8{1})

	nom2.MakeW(false, []uint8{7, 2})
	den2.MakeN([]uint8{1})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Subtraction(n1, n2)
	wantn.MakeW(true, []uint8{8})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("8 - получили %+v, ожидалось %+v", got, want)
	}
}

// Пименов
func TestMultiplication(t *testing.T) {
	var n1, n2, got, want Rational
	var nom1, nom2, wantn whole.Whole
	var den1, den2, wantd natural.Natural
	//////////////////////////////////////
	nom1.MakeW(false, []uint8{0})
	den1.MakeN([]uint8{1})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(false, []uint8{0})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}
	//////////////////////////////////////
	nom1.MakeW(false, []uint8{0})
	den1.MakeN([]uint8{3, 4})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(false, []uint8{0})
	wantd.MakeN([]uint8{1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}
	//////////////////////////////////////
	nom1.MakeW(false, []uint8{5})
	den1.MakeN([]uint8{5})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(false, []uint8{5})
	wantd.MakeN([]uint8{1, 9})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}
	//////////////////////////////////////
	nom1.MakeW(false, []uint8{1, 7})
	den1.MakeN([]uint8{1, 9})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(false, []uint8{8, 5})
	wantd.MakeN([]uint8{3, 6, 1})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("4 - получили %+v, ожидалось %+v", got, want)
	}
	//////////////////////////////////////
	nom1.MakeW(false, []uint8{4, 6})
	den1.MakeN([]uint8{5})

	nom2.MakeW(true, []uint8{8, 2})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(true, []uint8{3, 7, 7, 2})
	wantd.MakeN([]uint8{9, 5})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("5 - получили %+v, ожидалось %+v", got, want)
	}
	//////////////////////////////////////
	nom1.MakeW(true, []uint8{5, 6})
	den1.MakeN([]uint8{3})

	nom2.MakeW(true, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(false, []uint8{2, 8, 0})
	wantd.MakeN([]uint8{5, 7})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("6 - получили %+v, ожидалось %+v", got, want)
	}
	//////////////////////////////////////
	nom1.MakeW(true, []uint8{2})
	den1.MakeN([]uint8{2})

	nom2.MakeW(false, []uint8{5})
	den2.MakeN([]uint8{1, 9})

	n1.MakeR(nom1, den1)
	n2.MakeR(nom2, den2)

	got = Multiplication(n1, n2)
	wantn.MakeW(true, []uint8{5})
	wantd.MakeN([]uint8{1, 9})
	want.MakeR(wantn, wantd)
	if natural.Compare(wantd, got.Denominator) != 0 || natural.Compare(wantn.Num, got.Nominator.Num) != 0 || wantn.Negative != got.Nominator.Negative {
		t.Fatalf("7 - получили %+v, ожидалось %+v", got, want)
	}
}
