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
func Compare(a, b *Natural) int {
	var i uint32
	var flag, res int
	flag = 0
	switch {
	case a.Older > b.Older: //если в одном числе больше разрядов, чем в другом, то оно больше
		flag = 1
		res = 2
	case b.Older > a.Older:
		flag = 1
		res = 1
	default:
		for i = 0; i < a.Older; i++ { //сравниваем разряды чисел
			switch {
			case a.Digits[i] > b.Digits[i]:
				flag = 1
				res = 2
				break
			case b.Digits[i] > a.Digits[i]:
				flag = 1
				res = 1
				break
			}
		}
	}
	if flag == 0 {
		res = 0
	}
	return res
}

// CheckNull Турбина Проверка на ноль: если число не равно нулю, то «да» иначе «нет»
func CheckNull(x *Natural) bool {
	var res bool
	res = false
	if x.Older == 0 { //если в числе один разряд и он равен нулю, то число является нулем
		if x.Digits[0] == 0 {
			res = true
		}
	}
	return res
}
