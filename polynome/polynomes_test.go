package polynome

import (
	"DMTI_baza/rational"
	"fmt"
	"testing"
)

func TestAdditionP(t *testing.T) {
	var want, got, pol1, pol2 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
		},
	}
	pol2 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
		},
	}
	want = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(2, 1),
			rational.IntToRational(2, 1),
			rational.IntToRational(2, 1),
			rational.IntToRational(2, 1),
		},
	}
	got = AdditionP(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	pol1 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(-1, 1),
			rational.IntToRational(5, 1),
			rational.IntToRational(12, 1),
			rational.IntToRational(13, 1),
		},
	}
	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
		},
	}
	want = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(-1, 1),
			rational.IntToRational(5, 1),
			rational.IntToRational(13, 1),
			rational.IntToRational(14, 1),
		},
	}
	got = AdditionP(pol1, pol2)
	got = AdditionP(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestSubstractionP(t *testing.T) {
	var want, got, pol1, pol2 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
		},
	}
	pol2 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
		},
	}
	want = Polynomial{
		Older: 0,
		Coeff: []rational.Rational{
			rational.IntToRational(0, 1),
		},
	}
	got = SubstractionP(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, \nожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	pol1 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(-1, 1),
			rational.IntToRational(5, 1),
			rational.IntToRational(12, 1),
			rational.IntToRational(13, 1),
		},
	}
	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(0, 1),
		},
	}
	want = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(-1, 1),
			rational.IntToRational(5, 1),
			rational.IntToRational(11, 1),
			rational.IntToRational(13, 1),
		},
	}
	got = SubstractionP(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, \nожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestMultiplicationXpowerK(t *testing.T) {
	var want, got, pol1 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
		},
	}
	want = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(1, 1),
			rational.IntToRational(0, 1),
			rational.IntToRational(0, 1),
		},
	}
	got = MultiplicationXpowerK(pol1, 2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestMultiplicationRational(t *testing.T) {
	var want, got, pol1 Polynomial
	var rat rational.Rational
	no := 1
	pol1 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(2, 1),
			rational.IntToRational(4, 1),
		},
	}
	rat = rational.IntToRational(1, 2)
	want = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(2, 1),
		},
	}
	got = MultiplicationRational(pol1, rat)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, \nожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestQuotientOfDivision(t *testing.T) {
	var want, got, pol1, pol2 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(-5, 1),
			rational.IntToRational(-22, 1),
			rational.IntToRational(56, 1),
		},
	}

	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(4, 1),
		},
	}

	got = QuotientOfDivision(pol1, pol2)
	want = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(-9, 1),
			rational.IntToRational(14, 1),
		},
	}
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestGreatestCommonDivisor(t *testing.T) {
	var want, got, pol1, pol2 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(2, 1),
			rational.IntToRational(0, 1),
		},
	}

	pol2 = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(2, 1),
			rational.IntToRational(4, 1),
		},
	}

	got = GreatestCommonDivisor(pol1, pol2)
	want = Polynomial{
		Older: 0,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
		},
	}

	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	pol1 = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(4, 1),
			rational.IntToRational(4, 1),
		},
	}

	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(2, 1),
		},
	}

	got = GreatestCommonDivisor(pol1, pol2)
	want = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(2, 1),
		},
	}
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
}

func TestMultiplicationPol(t *testing.T) {
	var want, got, pol1, pol2 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(5, 1),
			rational.IntToRational(-6, 1),
		},
	}
	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(5, 1),
		},
	}
	want = Polynomial{
		Older: 3,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(10, 1),
			rational.IntToRational(19, 1),
			rational.IntToRational(-30, 1),
		},
	}
	got = MultiplicationPol(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestRemainderFromDivision(t *testing.T) {
	var want, got, pol1, pol2 Polynomial
	no := 1
	pol1 = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(6, 1),
			rational.IntToRational(8, 1),
		},
	}
	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(4, 1),
		},
	}
	want = Polynomial{
		Older: 0,
		Coeff: []rational.Rational{
			rational.IntToRational(0, 1),
		},
	}
	got = RemainderFromDivision(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	pol1 = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(6, 1),
			rational.IntToRational(8, 1),
		},
	}
	pol2 = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(5, 1),
		},
	}
	want = Polynomial{
		Older: 0,
		Coeff: []rational.Rational{
			rational.IntToRational(3, 1),
		},
	}
	got = RemainderFromDivision(pol1, pol2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestDerivative(t *testing.T) {
	var want, got, pol Polynomial
	no := 1
	pol = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(-1, 1),
			rational.IntToRational(2, 1),
			rational.IntToRational(0, 1),
		},
	}
	got = Derivative(pol)
	want = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(-2, 1),
			rational.IntToRational(2, 1),
		},
	}
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	pol = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(4, 1),
			rational.IntToRational(2, 1),
			rational.IntToRational(0, 1),
		},
	}
	got = Derivative(pol)
	want = Polynomial{
		Older: 1,
		Coeff: []rational.Rational{
			rational.IntToRational(8, 1),
			rational.IntToRational(2, 1),
		},
	}
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
}

func TestSimplifyRoots(t *testing.T) {
	var pol, got, want Polynomial
	no := 1

	pol = Polynomial{
		Older: 4,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(12, 1),
			rational.IntToRational(52, 1),
			rational.IntToRational(96, 1),
			rational.IntToRational(64, 1),
		},
	}

	got = SimplifyRoots(pol)
	want = Polynomial{
		Older: 2,
		Coeff: []rational.Rational{
			rational.IntToRational(1, 1),
			rational.IntToRational(6, 1),
			rational.IntToRational(8, 1),
		},
	}
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}
