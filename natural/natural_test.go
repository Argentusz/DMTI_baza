package natural

import "testing"

// Пименов
func TestSubtraction(t *testing.T) {
	var n1, n2 Natural
	n1.MakeN([]uint8{6, 4})
	n2.MakeN([]uint8{7, 2})
	got := Subtraction(n1, n2)
	want := []uint8{8}
	for i, v := range got.Digits {
		if v != want[i] {
			t.Fatalf("1 - получили %+v, ожидалось %+v", got, want)
		}
	}
}
