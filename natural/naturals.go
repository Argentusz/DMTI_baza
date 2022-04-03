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
		for i = 0; i < a.Older; i++ { //сравниваем разряды чисел, если соответствующий разряд больше, то одно число больше другого
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

// CheckNull Турбина Проверка на ноль: если число не равно нулю, то «да» иначе «нет»
func CheckNull(x Natural) bool {

	if x.Older == 0 { //если в числе один разряд и он равен нулю, то число является нулем
		if x.Digits[0] == 0 {
			return true
		}
	}
	return false
}

// Комаровский Subtraction Вычитание из первого большего натурального числа второго меньшего или равного( сделал за Милану)
func Subtraction(x1, x2 Natural) Natural {
	var a, b, res Natural
	var i, j, k int
	var mass []uint8
	// опрределяем большее число
	if Compare(x1, x2) == 0 {
		mass = append(mass, 0)
		res.MakeN(mass)
		return res
	} else if Compare(x1, x2) == 2 {
		a = x1
		b = x2
	} else if Compare(x1, x2) == 1 {
		a = x2
		b = x1
	}
	i = int(a.Older)
	j = int(b.Older)
	for j > -1 {
		switch {
		case a.Digits[i]-b.Digits[j] < 0:
			if a.Digits[i-1] != 0 {
				a.Digits[i-1] = a.Digits[i-1] - 1
			} else {
				// для случая если идет много нулей
				k = i
				for a.Digits[i-1] == 0 {
					k -= 1
					a.Digits[i] = 9
				}
				a.Digits[i-1] = a.Digits[i-1] - 1
			}
			mass = append([]uint8{a.Digits[i] + 10 - b.Digits[j]}, mass...)
		default:
			mass = append([]uint8{a.Digits[i] - b.Digits[j]}, mass...)
		}
		i -= 1
		j -= 1
	}
	for i > -1 {
		mass = append([]uint8{a.Digits[i]}, mass...)
		i -= 1
	}
	res.MakeN(mass)
	return res
}
