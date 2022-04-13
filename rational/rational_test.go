package rational

import (
	"DMTI_baza/natural"
	"DMTI_baza/whole"
	"fmt"
	"testing"
)

func TestIntToRational(t *testing.T) {
	var want, got Rational
	no := 1
	got = IntToRational(2, -5)
	want = Rational{
		Nominator: whole.Whole{
			Num: natural.Natural{
				Digits: []uint8{2},
				Older:  0,
			}, Negative: true,
		},
		Denominator: natural.Natural{
			Digits: []uint8{5},
			Older:  0,
		},
	}
	for i, v := range got.Nominator.Num.Digits {
		if v != want.Nominator.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Nominator.Negative != want.Nominator.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	for i, v := range got.Denominator.Digits {
		if v != want.Denominator.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++
}

func TestWholeToFractional(t *testing.T) {
	var w1 whole.Whole
	var got, want Rational
	no := 1
	w1 = whole.IntToWhole(-12)
	want = IntToRational(-12, 1)
	got = WholeToFractional(w1)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = whole.IntToWhole(12)
	want = IntToRational(12, 1)
	got = WholeToFractional(w1)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestFractionalToWhole(t *testing.T) {
	var r1 Rational
	var got, want whole.Whole
	no := 1
	r1 = IntToRational(12, 1)
	want = whole.IntToWhole(12)
	got = FractionalToWhole(r1)
	if whole.Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(-12, 1)
	want = whole.IntToWhole(-12)
	got = FractionalToWhole(r1)
	if whole.Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
}

func TestCheckingForWhole(t *testing.T) {
	var r1 Rational
	var got, want bool
	no := 1
	r1 = IntToRational(12, 5)
	want = false
	got = CheckingForWhole(r1)
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(12, 1)
	want = true
	got = CheckingForWhole(r1)
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
}

func TestSimplifyingFractions(t *testing.T) {
	var r1, want, got Rational
	no := 1
	r1 = IntToRational(12, 6)
	want = IntToRational(2, 1)
	got = SimplifyingFractions(r1)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(12, 10)
	want = IntToRational(6, 5)
	got = SimplifyingFractions(r1)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(12, 7)
	want = IntToRational(12, 7)
	got = SimplifyingFractions(r1)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestMultiplication(t *testing.T) {
	var r1, r2, want, got Rational
	no := 1
	r1 = IntToRational(2, 1)
	r2 = IntToRational(3, 1)
	want = IntToRational(6, 1)
	got = Multiplication(r1, r2)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(2, 1)
	r2 = IntToRational(0, 1)
	want = IntToRational(0, 1)
	got = Multiplication(r1, r2)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(3, 5)
	r2 = IntToRational(5, 3)
	want = IntToRational(1, 1)
	got = Multiplication(r1, r2)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestDivision(t *testing.T) {
	var r1, r2, want, got Rational
	no := 1
	r1 = IntToRational(6, 1)
	r2 = IntToRational(2, 1)
	want = IntToRational(3, 1)
	got = Division(r1, r2)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	r1 = IntToRational(3, 5)
	r2 = IntToRational(3, 5)
	want = IntToRational(1, 1)
	got = Division(r1, r2)
	if Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestSubtraction(t *testing.T) {
	var r1, r2 Rational
	var denom1, denom2, denomw natural.Natural
	var nom1, nom2, nomw whole.Whole
	var want, got Rational
	no := 1

	nom1.MakeW(false, []uint8{1})
	denom1.MakeN([]uint8{7})
	nom2.MakeW(false, []uint8{8})
	denom2.MakeN([]uint8{2, 9})
	nomw.MakeW(true, []uint8{2, 7})
	denomw.MakeN([]uint8{2, 0, 3})
	r1.MakeR(nom1, denom1)
	r2.MakeR(nom2, denom2)
	want.MakeR(nomw, denomw)
	fmt.Println(r1, "-", r2, "=", want)
	got = Subtraction(r1, r2)
	fmt.Println(got)
	for i, v := range got.Nominator.Num.Digits {
		if v != want.Nominator.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Nominator.Negative != want.Nominator.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	for i, v := range got.Denominator.Digits {
		if v != want.Denominator.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++
}

func TestAddition(t *testing.T) {
	var r1, r2 Rational
	var denom1, denom2, denomw natural.Natural
	var nom1, nom2, nomw whole.Whole
	var want, got Rational
	no := 1
	nom1.MakeW(false, []uint8{1})
	denom1.MakeN([]uint8{7})
	nom2.MakeW(false, []uint8{8})
	denom2.MakeN([]uint8{2, 9})
	nomw.MakeW(false, []uint8{8, 5})
	denomw.MakeN([]uint8{2, 0, 3})
	r1.MakeR(nom1, denom1)
	r2.MakeR(nom2, denom2)
	want.MakeR(nomw, denomw)
	fmt.Println(r1, "+", r2, "=", want)
	got = Addition(r1, r2)
	fmt.Println(got)
	for i, v := range got.Nominator.Num.Digits {
		if v != want.Nominator.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Nominator.Negative != want.Nominator.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	for i, v := range got.Denominator.Digits {
		if v != want.Denominator.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++

	nom1.MakeW(false, []uint8{8})
	denom1.MakeN([]uint8{2, 9})
	nom2.MakeW(false, []uint8{0})
	denom2.MakeN([]uint8{1})
	nomw.MakeW(false, []uint8{8})
	denomw.MakeN([]uint8{2, 9})
	r1.MakeR(nom1, denom1)
	r2.MakeR(nom2, denom2)
	want.MakeR(nomw, denomw)
	fmt.Println(r1, "+", r2, "=", want)
	got = Addition(r1, r2)
	fmt.Println(got)
	for i, v := range got.Nominator.Num.Digits {
		if v != want.Nominator.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Nominator.Negative != want.Nominator.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	for i, v := range got.Denominator.Digits {
		if v != want.Denominator.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++

	nom1.MakeW(false, []uint8{1, 4, 5})
	denom1.MakeN([]uint8{1, 7, 7})
	nom2.MakeW(true, []uint8{9, 1, 3})
	denom2.MakeN([]uint8{7})
	nomw.MakeW(true, []uint8{1, 6, 0, 5, 8, 6})
	denomw.MakeN([]uint8{1, 2, 3, 9})
	r1.MakeR(nom1, denom1)
	r2.MakeR(nom2, denom2)
	want.MakeR(nomw, denomw)
	fmt.Println(r1, "+", r2, "=", want)
	got = Addition(r1, r2)
	fmt.Println(got)
	for i, v := range got.Nominator.Num.Digits {
		if v != want.Nominator.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Nominator.Negative != want.Nominator.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	for i, v := range got.Denominator.Digits {
		if v != want.Denominator.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++
}
