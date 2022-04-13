package whole

import "testing"

func TestAddition(t *testing.T) {
	var n1, n2 Whole
	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{5, 1, 4})
	got := Addition(n1, n2)
	want := []uint8{1}
	wantb := true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("1 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 4})
	n2.MakeW(true, []uint8{5, 1, 3})
	got = Addition(n1, n2)
	want = []uint8{1, 0, 2, 7}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("2 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{5, 1, 4})
	got = Addition(n1, n2)
	want = []uint8{1, 0, 2, 7}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("3 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(false, []uint8{5, 1, 4})
	got = Addition(n1, n2)
	want = []uint8{1, 0, 2, 7}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("4 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{4, 8, 7})
	n2.MakeW(false, []uint8{1, 3})
	got = Addition(n1, n2)
	want = []uint8{5, 0, 0}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("5 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{4, 8, 7})
	n2.MakeW(true, []uint8{1, 3})
	got = Addition(n1, n2)
	want = []uint8{5, 0, 0}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("6 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{5, 0, 0})
	got = Addition(n1, n2)
	want = []uint8{1, 3}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("7 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 3})
	n2.MakeW(false, []uint8{5, 0, 0})
	got = Addition(n1, n2)
	want = []uint8{1, 3}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("8 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 0, 0})
	n2.MakeW(false, []uint8{8, 0, 0})
	got = Addition(n1, n2)
	want = []uint8{1, 3, 0, 0}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("9 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 0, 0})
	n2.MakeW(true, []uint8{8, 0, 0})
	got = Addition(n1, n2)
	want = []uint8{1, 3, 0, 0}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("10 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
}
func TestMultiplicationByNegativeOne(t *testing.T) {
	var n Whole
	n.MakeW(false, []uint8{8})
	got := MultiplicationByNegativeOne(n)
	want := []uint8{8}
	wantb := true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("1 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
}

func TestSubtraction(t *testing.T) {

	var n1, n2 Whole

	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{5, 1, 4})
	got := Subtraction(n1, n2)
	want := []uint8{1, 0, 2, 7}
	wantb := false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("1 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 4})
	n2.MakeW(true, []uint8{5, 1, 3})
	got = Subtraction(n1, n2)
	want = []uint8{1}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("2 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{5, 1, 4})
	got = Subtraction(n1, n2)
	want = []uint8{1}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("3 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(false, []uint8{5, 1, 4})
	got = Subtraction(n1, n2)
	want = []uint8{1}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("4 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(false, []uint8{1, 3})
	got = Subtraction(n1, n2)
	want = []uint8{5, 0, 0}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("5 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{1, 3})
	got = Subtraction(n1, n2)
	want = []uint8{5, 0, 0}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("6 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 1, 3})
	n2.MakeW(false, []uint8{5, 0, 0})
	got = Subtraction(n1, n2)
	want = []uint8{1, 3}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("7 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 1, 3})
	n2.MakeW(true, []uint8{5, 0, 0})
	got = Subtraction(n1, n2)
	want = []uint8{1, 3}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("8 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(false, []uint8{5, 0, 0})
	n2.MakeW(false, []uint8{8, 0, 0})
	got = Subtraction(n1, n2)
	want = []uint8{3, 0, 0}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("9 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
	n1.MakeW(true, []uint8{5, 0, 0})
	n2.MakeW(true, []uint8{8, 0, 0})
	got = Subtraction(n1, n2)
	want = []uint8{3, 0, 0}
	wantb = false
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("10 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}

	n1.MakeW(false, []uint8{6, 4})
	n2.MakeW(false, []uint8{7, 2})
	got = Subtraction(n1, n2)
	want = []uint8{8}
	wantb = true
	for i, v := range got.Num.Digits {
		if v != want[i] || got.Negative != wantb {
			t.Fatalf("11 - получили %+v, ожидалось %+v, %+v", got, want, wantb)
		}
	}
}
