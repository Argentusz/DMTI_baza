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

// Compare Турбина Сравнение натуральных чисел: 2 - если первое больше второго, 0, если равно, 1 иначе.
func Compare(a, b Natural) int {
	var i uint32

	switch {
	case a.Older > b.Older: //если в одном числе больше разрядов, чем в другом, то оно больше
		return 2
	case b.Older > a.Older:
		return 1
	default:
		for i = 0; i < a.Older+1; i++ { //сравниваем разряды чисел, если соответствующий разряд больше, то одно число больше другого
			switch {
			case a.Digits[i] > b.Digits[i]:
				return 2
			case b.Digits[i] > a.Digits[i]:
				return 1
			}
		}
	}

	return 0
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
func Addition(a, b Natural) Natural {
	var i uint32
	var r, t, buffer Natural
	r = CopyN(a)
	t = CopyN(b)
	if Compare(r, t) != 2 && Compare(r, t) != 0 { //Сравниваем числа, если первое небольше второго и они оба не равны, то меняем их местами
		buffer = r
		r = t
		t = buffer
	}
	for i = 0; i <= t.Older; i++ { //Цикл прибавления последней цифры одного числа к другой, смещаемся влево до тех пор, пока не дойдём до конца меньшего
		r.Digits[r.Older-i] += t.Digits[t.Older-i]
	}
	for i = 0; i <= r.Older; i++ { //теперь проходим по получившемуся числу и, если где-то остался элемент больше или равный 10, исправляем
		if r.Digits[r.Older-i] >= 10 {
			if r.Older-i == 0 && r.Digits[0] >= 10 { //если очередь дошла до последнего разряда(самый левый) и если он  больше или равен 10, то
				r.Digits = append([]uint8{0}, r.Digits...) //добавляю в начало числа 0
				r.Older += 1                               //увеличиваю older ("размер" числа?)
				r.Digits[0] = r.Digits[1] / 10
				r.Digits[1] %= 10
			} else { // если нет, то вычитаем из тек разряда 10 и +1 к след
				r.Digits[r.Older-i] -= 10
				r.Digits[r.Older-i-1] += 1
			}
		}
	}
	return r
}
