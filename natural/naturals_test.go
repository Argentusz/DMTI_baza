package natural

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	var n1, n2 Natural
	var got, want int
	no := 1
	n1 = IntToNat(12)
	n2 = IntToNat(13)
	got = Compare(n1, n2)
	want = 1
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(14)
	n2 = IntToNat(13)
	got = Compare(n1, n2)
	want = 2
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(12)
	n2 = IntToNat(12)
	got = Compare(n1, n2)
	want = 0
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestCheckNull(t *testing.T) {
	var n1 Natural
	var want, got bool
	no := 1
	n1 = IntToNat(10)
	got = CheckNull(n1)
	want = false
	if got != want {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(0)
	got = CheckNull(n1)
	want = true
	if got != want {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestAddition1(t *testing.T) {
	var n1, got, want Natural
	no := 1
	n1 = IntToNat(0)
	got = Addition1(n1)
	want = IntToNat(1)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	n1 = IntToNat(123)
	got = Addition1(n1)
	want = IntToNat(124)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestMultiplicationNaturalNumber(t *testing.T) {
	var n1, got, want Natural
	var digit uint8
	no := 1
	n1 = IntToNat(10)
	digit = 2
	want = IntToNat(20)
	got = MultiplicationNaturalNumber(n1, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(0)
	digit = 9
	want = IntToNat(0)
	got = MultiplicationNaturalNumber(n1, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(99)
	digit = 2
	want = IntToNat(198)
	got = MultiplicationNaturalNumber(n1, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestMultiplicationBy10k(t *testing.T) {
	var n1, got, want Natural
	var digit int
	no := 1
	n1 = IntToNat(10)
	digit = 2
	want = IntToNat(1000)
	got = MultiplicationBy10k(n1, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(0)
	digit = 9
	want = IntToNat(0)
	got = MultiplicationBy10k(n1, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(99)
	digit = 2
	want = IntToNat(9900)
	got = MultiplicationBy10k(n1, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestDivideOneIteration(t *testing.T) {
	var n1, n2, want, got Natural
	no := 1
	n1 = IntToNat(123)
	n2 = IntToNat(3)
	want = IntToNat(40)
	got = DivideOneIteration(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestAddition(t *testing.T) {
	var n1, n2, want, got Natural
	no := 1
	n1 = IntToNat(123)
	n2 = IntToNat(3)
	want = IntToNat(126)
	got = Addition(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(0)
	n2 = IntToNat(3)
	want = IntToNat(3)
	got = Addition(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestSubtraction(t *testing.T) {
	var n1, n2, want, got Natural
	no := 1
	n1 = IntToNat(123)
	n2 = IntToNat(3)
	want = IntToNat(120)
	got = Subtraction(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(3)
	n2 = IntToNat(0)
	want = IntToNat(3)
	got = Subtraction(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestDifferenceOfNaturals(t *testing.T) {
	var n1, n2, want, got Natural
	var digit uint8
	no := 1
	n1 = IntToNat(123)
	n2 = IntToNat(3)
	want = IntToNat(117)
	digit = 2
	got = DifferenceOfNaturals(n1, n2, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(3)
	n2 = IntToNat(0)
	digit = 123
	want = IntToNat(3)
	got = DifferenceOfNaturals(n1, n2, digit)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestMultiplication(t *testing.T) {
	var n1, n2, got, want Natural
	no := 1
	n1 = IntToNat(10)
	n2 = IntToNat(123)
	want = IntToNat(1230)
	got = Multiplication(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(0)
	n2 = IntToNat(123)
	want = IntToNat(0)
	got = Multiplication(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestRemainderFromDivision(t *testing.T) {
	var n1, n2, got, want Natural
	no := 1
	n1 = IntToNat(123)
	n2 = IntToNat(5)
	want = IntToNat(3)
	got = RemainderFromDivision(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(125)
	n2 = IntToNat(5)
	want = IntToNat(0)
	got = RemainderFromDivision(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestIntegerFromDivision(t *testing.T) {
	var n1, n2, want, got Natural
	no := 1
	n1.MakeN([]uint8{1, 3, 1, 1})
	n2.MakeN([]uint8{1, 3})

	got = IntegerFromDivision(n1, n2)
	want.MakeN([]uint8{1, 0, 0})
	for i, v := range got.Digits {
		if v != want.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++
}

func TestGreatestCommonDivisor(t *testing.T) {
	var n1, n2, want, got Natural
	no := 1
	n1.MakeN([]uint8{1, 3, 3})
	n2.MakeN([]uint8{1, 3})

	got = GreatestCommonDivisor(n1, n2)
	want.MakeN([]uint8{1})

	for i, v := range got.Digits {
		if v != want.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	fmt.Println(no, "Passed")
	no++
}

func TestLeastCommonMultiple(t *testing.T) {
	var n1, n2, got, want Natural
	no := 1
	n1 = IntToNat(2)
	n2 = IntToNat(5)
	want = IntToNat(10)
	got = LeastCommonMultiple(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	n1 = IntToNat(2)
	n2 = IntToNat(4)
	want = IntToNat(4)
	got = LeastCommonMultiple(n1, n2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}
