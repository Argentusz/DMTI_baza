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
