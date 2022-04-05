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

//Хвостовский Добавление 1 к натуральному числу
func Addition1(x Natural) Natural {
	var i int
	if x.Digits[x.Older] == 9 { // Если последний разряд равен 9, то при прибавлении единицы, он становится 0
		x.Digits[x.Older] = 0
		for i = int(x.Older - 1); i >= 0; i-- { // Прибавляем единицу к первому разряду, не равному 9(9 заменяем на 0)
			if x.Digits[i] < 9 {
				x.Digits[i] += 1
				return x
			} else {
				x.Digits[i] = 0
			}
		}
	} else {
		x.Digits[x.Older] += 1 //Если последний разряд не равен 9, то просто к нему прибавляем единицу
		return x
	}
	x.Digits[0] = 1                // Если все разряды равны 9, то результатов будет число с первой единицей, а остальными нулями
	x.Older += 1                   // На разряд больше, чем изначальное
	x.Digits = append(x.Digits, 0) // Присваиваем 0 только последнему разряду, так как остальным он уже присвоен в цикле
	return x
}

// MultiplicationNaturalNumber Хвостовский Умножение натурального числа на цифру
func MultiplicationNaturalNumber(x Natural, b uint8) Natural {
	var c uint8
	var i int
	c = 0
	for i = int(x.Older); i >= 0; i-- { //берём последний элемент массива до первого элемента массива
		x.Digits[i] *= b       // умножаю разряд числа на цифру
		x.Digits[i] += c       // перенос лишнего десятка с умножения
		if x.Digits[i] >= 10 { //если текущий разряд больше или равен 10
			c = x.Digits[i] / 10 // записываем в переменную лишние десятки
			x.Digits[i] = x.Digits[i] % 10
		} else {
			c = 0 // если лишних десятков нет, то обнуляем
		}
	}
	if c != 0 { // если остался лишний десяток, то...
		x.Digits = append(x.Digits, x.Digits[x.Older]) //увеличиваем массив
		x.Older += 1
		for i = int(x.Older - 1); i > 0; i-- { // смещаем массив вправо
			x.Digits[i] = x.Digits[i-1]
		}
		x.Digits[0] = c // первый элемент массива приравниваем к остатку
	}
	return x
}

// MultiplicationBy10k Хвостовский Умножение натурального числа на 10^k
func MultiplicationBy10k(x Natural, k int) Natural {
	var i int
	if k == 0 { //если степень равна нулю, то возвращаем изначальное значение числа
		return x
	} else { //иначе увеличваем число на k разрядов
		for i = 0; i < k; i++ {
			x.Digits = append(x.Digits, 0)
			x.Older += 1
		}
		return x
	}
}
