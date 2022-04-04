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

// Семёнов Addiction Сложение двух наутральных чисел
func Addiction(a, b *Natural) Natural {
	var i, j uint32
	var buf1, buf2 uint8
	if Compare(a, b) == 2 || Compare(a, b) == 0 {
		for i = 0; i <= b.Older; i++ {
			a.Digits[a.Older-i] += b.Digits[b.Older-i]
			if a.Digits[a.Older-i] >= 10 {
				if a.Digits[0] >= 10 {
					a.Digits = append(a.Digits, a.Digits[a.Older])
					a.Older += 1
					buf2 = a.Digits[0]
					a.Digits[0] /= 10
					buf1 = a.Digits[1]
					a.Digits[1] = buf2 % 10

					for j = 2; j < a.Older; j++ {
						buf2 = a.Digits[j]
						a.Digits[j] = buf1
						buf1 = buf2
					}
				} else {
					a.Digits[a.Older-i] -= 10
					a.Digits[a.Older-i-1] += 1
				}
			}
		}
		return *a
	} else {
		for i = 0; i <= a.Older; i++ {
			b.Digits[b.Older-i] += a.Digits[a.Older-i]
			if b.Digits[b.Older-i] >= 10 {
				if b.Digits[0] >= 10 {
					b.Digits = append(b.Digits, b.Digits[b.Older])
					b.Older += 1
					buf2 = b.Digits[0]
					b.Digits[0] /= 10
					buf1 = b.Digits[1]
					b.Digits[1] = buf2 % 10

					for j = 2; j < b.Older; j++ {
						buf2 = b.Digits[j]
						b.Digits[j] = buf1
						buf1 = buf2
					}
				} else {
					b.Digits[b.Older-i] -= 10
					b.Digits[b.Older-i-1] += 1
				}
			}
		}
		return *b
	}
}
