package natural

// Natural Структура натурального числа
type Natural struct {
	Digits []uint8 // Цифры больше 255 и меньше 0 нам не понадобятся, а памяти сохранит уйму
	Older  uint32
}

// MakeN Метод для создания натурала
func (n *Natural) MakeN(digits []uint8) {
	for _, v := range digits {
		if v != 0 { //Проверка на 0 старшей цифры
			break
		} else {
			digits = digits[1:]
		}
	}
	n.Digits = digits
	n.Older = uint32(len(digits)) - 1
}

// Турбина Compare Сравнение натуральных чисел: 2 - если первое больше второго, 0, если равно, 1 иначе.
func Compare(a, b *Natural) int {
	var i uint32
	var flag, res int
	flag = 0
	switch {
	case a.Older > b.Older:
		flag = 1
		res = 2
	case b.Older > a.Older:
		flag = 1
		res = 1
	default:
		for i = 0; i < a.Older; i++ {
			switch {
			case a.Digits[0] > b.Digits[0]:
				flag = 1
				res = 2
			case b.Digits[0] > a.Digits[0]:
				flag = 1
				res = 1
			}
		}
	}
	if flag == 0 {
		res = 0
	}
	return res
}

//Тростин Функция для копирования натурального числа
func CopyN(n Natural) Natural {
	var i uint32
	var x Natural
	for i = 0; i <= n.Older; i++ {
		x.Digits = append(x.Digits, n.Digits[i])
	}
	x.Older = n.Older
	return x
}

// Семёнов Addition Сложение двух наутральных чисел
func Addition(a, b *Natural) Natural {
	var i uint32
	var r, t, d Natural
	r = CopyN(*a)
	t = CopyN(*b)
	y := &r
	u := &t
	if Compare(y, u) != 2 || Compare(y, u) != 0 {
		d = r
		r = t
		t = d
	}
	for i = 0; i <= t.Older; i++ { //Цикл прибавления последней цифры одного числа к другой, смещаемся влево до тех пор, пока не дойдём до конца меньшего
		r.Digits[r.Older-i] += t.Digits[t.Older-i]
		if r.Digits[r.Older-i] >= 10 { //если текущий разряд больше или равен 10
			if r.Digits[0] >= 10 { //если именно последний разряд(самый левый) больше или равен 10
				r.Digits = append([]uint8{0}, r.Digits...) //добавляю в начало числа 0
				r.Older += 1                               //увеличиваю older ("размер" числа?)
				r.Digits[0] = r.Digits[1] / 10
				r.Digits[1] %= 10

			} else { //если последний разряд(самый левый) не больше или равен 10, то просто вычитаю из текущего 10 и прибавляю к след 1
				r.Digits[a.Older-i] -= 10
				r.Digits[a.Older-i-1] += 1
			}
		}
	}
	return r
}
