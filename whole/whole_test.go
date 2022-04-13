package whole

import (
	"DMTI_baza/natural"
	"fmt"
	"testing"
)

func TestAbsolute(t *testing.T) {
	var w1 Whole
	var got, want natural.Natural
	no := 1
	w1 = IntToWhole(-1)
	want = natural.IntToNat(1)
	got = Absolute(w1)
	if natural.Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(1)
	want = natural.IntToNat(1)
	got = Absolute(w1)
	if natural.Compare(want, got) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestPositivity(t *testing.T) {
	var w1 Whole
	var got, want int
	no := 1
	w1 = IntToWhole(-12)
	want = 1
	got = Positivity(w1)
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(12)
	want = 2
	got = Positivity(w1)
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(0)
	want = 0
	got = Positivity(w1)
	if want != got {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestMultiplicationByNegativeOne(t *testing.T) {
	var w1, got, want Whole
	no := 1
	w1 = IntToWhole(-12)
	want = IntToWhole(12)
	got = MultiplicationByNegativeOne(w1)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(12)
	want = IntToWhole(-12)
	got = MultiplicationByNegativeOne(w1)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(0)
	want = IntToWhole(0)
	got = MultiplicationByNegativeOne(w1)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestFromNaturalsToWhole(t *testing.T) {
	var n1 natural.Natural
	var got, want Whole
	no := 1
	n1 = natural.IntToNat(12)
	want = IntToWhole(12)
	got = FromNaturalsToWhole(n1)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestFromWholeToNaturals(t *testing.T) {
	var w1 Whole
	var got, want natural.Natural
	no := 1
	w1 = IntToWhole(12)
	want = natural.IntToNat(12)
	got = FromWholeToNaturals(w1)
	if natural.Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestIntToWhole(t *testing.T) {
	var got, want Whole
	var num int64
	no := 1

	num = 123
	got = IntToWhole(num)
	want.MakeW(false, []uint8{1, 2, 3})
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	num = 0
	got = IntToWhole(num)
	want.MakeW(false, []uint8{0})
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	num = -123
	got = IntToWhole(num)
	want.MakeW(true, []uint8{1, 2, 3})
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestWholeFromDivision(t *testing.T) {
	var w1, w2 Whole
	no := 1
	w1.MakeW(false, []uint8{1, 2})
	w2.MakeW(false, []uint8{1})
	got := WholeFromDivision(w1, w2)
	want := Whole{Num: natural.Natural{Digits: []uint8{1, 2}, Older: 1}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(false, []uint8{1, 7})
	w2.MakeW(true, []uint8{5})
	got = WholeFromDivision(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{4}, Older: 0}, Negative: true}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(true, []uint8{4})
	w2.MakeW(true, []uint8{2})
	got = WholeFromDivision(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{2}, Older: 0}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestRemainderFromDivision(t *testing.T) {
	var w1, w2 Whole
	no := 1
	w1.MakeW(false, []uint8{1, 2})
	w2.MakeW(false, []uint8{1})
	got := RemainderFromDivision(w1, w2)
	want := Whole{Num: natural.Natural{Digits: []uint8{0}, Older: 0}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(false, []uint8{1, 7})
	w2.MakeW(true, []uint8{5})
	got = RemainderFromDivision(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{3}, Older: 0}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

}

func TestAddition(t *testing.T) {
	var w1, w2, got, want Whole
	no := 1
	w1 = IntToWhole(-12)
	w2 = IntToWhole(12)
	want = Zero()
	got = Addition(w1, w2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(13)
	w2 = IntToWhole(12)
	want = IntToWhole(25)
	got = Addition(w1, w2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
	w1 = IntToWhole(-1)
	w2 = IntToWhole(12)
	want = IntToWhole(11)
	got = Addition(w1, w2)
	if Compare(got, want) != 0 {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestSubtraction(t *testing.T) {
	var w1, w2 Whole
	no := 1
	w1.MakeW(false, []uint8{1, 2})
	w2.MakeW(false, []uint8{1})
	got := Subtraction(w1, w2)
	want := Whole{Num: natural.Natural{Digits: []uint8{1, 1}, Older: 1}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(false, []uint8{1, 2})
	w2.MakeW(true, []uint8{1})
	got = Subtraction(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{1, 3}, Older: 1}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(false, []uint8{1, 2})
	w2.MakeW(false, []uint8{1, 5})
	got = Subtraction(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{3}, Older: 0}, Negative: true}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(true, []uint8{1, 2})
	w2.MakeW(false, []uint8{1, 5})
	got = Subtraction(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{2, 7}, Older: 0}, Negative: true}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}

func TestMultiplication(t *testing.T) {
	var w1, w2 Whole
	no := 1
	w1.MakeW(false, []uint8{1, 2})
	w2.MakeW(false, []uint8{1})
	got := Multiplication(w1, w2)
	want := Whole{Num: natural.Natural{Digits: []uint8{1, 2}, Older: 1}, Negative: false}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(true, []uint8{1, 2})
	w2.MakeW(false, []uint8{1})
	got = Multiplication(w1, w2)
	want = Whole{Num: natural.Natural{Digits: []uint8{1, 2}, Older: 1}, Negative: true}
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(true, []uint8{1, 2})
	w2.MakeW(true, []uint8{1})
	got = Multiplication(w1, w2)
	want.MakeW(false, []uint8{1, 2})
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(true, []uint8{7, 2})
	w2.MakeW(false, []uint8{1, 4, 8, 8})
	got = Multiplication(w1, w2)
	want.MakeW(true, []uint8{1, 0, 7, 1, 3, 6})
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++

	w1.MakeW(false, []uint8{0})
	w2.MakeW(true, []uint8{1, 4, 8, 8})
	got = Multiplication(w1, w2)
	want.MakeW(false, []uint8{0})
	for i, v := range got.Num.Digits {
		if v != want.Num.Digits[i] {
			t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
		}
	}
	if got.Negative != want.Negative {
		t.Fatalf("%d - получили %+v, ожидалось %+v", no, got, want)
	}
	fmt.Println(no, "Passed")
	no++
}
