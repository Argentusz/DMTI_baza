package natural

import (
	"testing"
)

//тестовые функции Голубева Михаила
//тест функции сравнения натуральных чисел
func TestCompare(t *testing.T) {
	var n1, n2 Natural
	n1.Digits = []uint8{1, 2, 3, 4, 5, 6, 7}
	n1.Older = 6
	n2.Digits = []uint8{1, 2, 3, 4, 5, 6, 7}
	n2.Older = 6
	got := Compare(n1, n2)
	want := 0
	if got != want {
		t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
	}

	n1.Digits = []uint8{1, 2, 3, 4, 5, 6, 6}
	n1.Older = 6
	n2.Digits = []uint8{1, 2, 3, 4, 5, 6, 7}
	n2.Older = 6
	got = Compare(n1, n2)
	want = 1
	if got != want {
		t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
	}

	n1.Digits = []uint8{1, 2, 3, 4, 5, 6, 7}
	n1.Older = 6
	n2.Digits = []uint8{1, 2, 3, 4, 5, 6, 6}
	n2.Older = 6
	got = Compare(n1, n2)
	want = 2
	if got != want {
		t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
	}
}

//тест добавления единички
func TestAddition1(t *testing.T) {
	var n Natural
	n.Digits = []uint8{1, 2, 3, 4, 5, 6, 7}
	n.Older = 6
	got := Addition1(n)
	want := []uint8{1, 2, 3, 4, 5, 6, 8}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n.Digits = []uint8{1, 2, 3, 4, 5, 6, 0}
	n.Older = 6
	got = Addition1(n)
	want = []uint8{1, 2, 3, 4, 5, 6, 1}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("2 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n.Digits = []uint8{1, 2, 3, 4, 5, 6, 9}
	n.Older = 6
	got = Addition1(n)
	want = []uint8{1, 2, 3, 4, 5, 7, 0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("3 - получили %+v, ожидалось %+v", got, want)
		}
	}
}

//тест умножения
func TestMultiplication(t *testing.T) {
	var n1, n2 Natural
	n1.MakeN([]uint8{4})
	n2.MakeN([]uint8{1})
	got := Multiplication(n1, n2)
	want := []uint8{4}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{4})
	n2.MakeN([]uint8{0})
	got = Multiplication(n1, n2)
	want = []uint8{0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{3})
	n2.MakeN([]uint8{9})
	got = Multiplication(n1, n2)
	want = []uint8{2, 7}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{0})
	n2.MakeN([]uint8{1, 2, 4, 5})
	got = Multiplication(n1, n2)
	want = []uint8{0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{3, 7, 8})
	n2.MakeN([]uint8{1, 2, 4, 5})
	got = Multiplication(n1, n2)
	want = []uint8{4, 7, 0, 6, 1, 0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
}

func TestIntegerFromDivision(t *testing.T) {
	var n1, n2 Natural
	n1.MakeN([]uint8{5, 1, 2})
	n2.MakeN([]uint8{2})
	got := IntegerFromDivision(n1, n2)
	want := []uint8{2, 5, 6}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{5, 1, 3})
	n2.MakeN([]uint8{2})
	got = IntegerFromDivision(n1, n2)
	want = []uint8{2, 5, 6}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{5, 1, 3})
	n2.MakeN([]uint8{5, 1, 3})
	got = IntegerFromDivision(n1, n2)
	want = []uint8{1}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{2})
	n2.MakeN([]uint8{5, 1, 3})
	got = IntegerFromDivision(n1, n2)
	want = []uint8{2, 5, 6}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
}

func TestRemainderFromDivision(t *testing.T) {
	var n1, n2 Natural
	n1.MakeN([]uint8{5, 1, 2})
	n2.MakeN([]uint8{2})
	got := RemainderFromDivision(n1, n2)
	want := []uint8{0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{5, 1, 3})
	n2.MakeN([]uint8{2})
	got = RemainderFromDivision(n1, n2)
	want = []uint8{1}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{5, 1, 3})
	n2.MakeN([]uint8{5, 1, 3})
	got = RemainderFromDivision(n1, n2)
	want = []uint8{0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{2, 3})
	n2.MakeN([]uint8{1, 3})
	got = RemainderFromDivision(n1, n2)
	want = []uint8{1, 0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
	n1.MakeN([]uint8{1, 3})
	n2.MakeN([]uint8{2, 3})
	got = RemainderFromDivision(n1, n2)
	want = []uint8{1, 0}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
}
