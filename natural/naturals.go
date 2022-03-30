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
