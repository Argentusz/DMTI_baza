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
	n.Older = uint32(len(digits))
}
