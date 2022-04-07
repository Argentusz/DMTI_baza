package natural

import "fmt"

// Natural
// Структура натурального числа
type Natural struct {
	Digits []uint8 // Цифры больше 255 и меньше 0 нам не понадобятся, а памяти сохранит уйму
	Older  uint32
}

// Zero
// Возвращает натуральный нуль
func Zero() Natural {
	var zero Natural
	zero.Older = 0
	zero.Digits = append(zero.Digits, 0)
	return zero
}

// ToStringN
// Возвращает натуральное как строку
func ToStringN(n Natural) string {
	var s string
	for _, v := range n.Digits {
		s += fmt.Sprint(v)
	}
	return s
}

// MakeN
// Метод для создания натурала
func (n *Natural) MakeN(digits []uint8) {
	for _, v := range digits {
		if v != 0 { //Проверка на 0 старшей цифры
			break
		} else {
			digits = digits[1:]
		}
	}
	if len(digits) == 0 {
		digits = append(digits, 0)
	}
	n.Digits = digits
	n.Older = uint32(len(digits)) - 1
}

// CopyN Тростин Максим
// Функция для копирования натурального числа
func CopyN(n Natural) Natural {
	var i uint32
	var x Natural
	for i = 0; i <= n.Older; i++ {
		x.Digits = append(x.Digits, n.Digits[i])
	}
	x.Older = n.Older
	return x
}

// Compare Турбина
// Сравнение натуральных чисел: 2 - если первое больше второго, 0, если равно, 1 иначе.
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

// CheckNull Турбина
// Проверка на ноль: если число не равно нулю, то «да» иначе «нет»
func CheckNull(x Natural) bool {

	if x.Older == 0 { //если в числе один разряд и он равен нулю, то число является нулем
		if x.Digits[0] == 0 {
			return true
		}
	}
	return false
}

// Addition1 Хвостовский
// Добавление 1 к натуральному числу
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

// MultiplicationNaturalNumber Хвостовский
// Умножение натурального числа на цифру
func MultiplicationNaturalNumber(g Natural, b uint8) Natural {
	var c uint8
	var i int
	c = 0
	x := CopyN(g)
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

// MultiplicationBy10k Хвостовский
// Умножение натурального числа на 10^k
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

func DivideOneIteration(x, y Natural) Natural {
	var big, small, res Natural
	var multiplier uint8
	if Compare(x, y) == 0 {
		res.MakeN([]uint8{1})
		return res
	}
	if Compare(x, y) == 2 {
		big = CopyN(x)
		small = CopyN(y)
	} else {
		big = CopyN(y)
		small = CopyN(x)
	}

	k := big.Older - small.Older
	digitsOfSmaller := small.Older + 1
	necessaryBig := CopyN(big)
	necessaryBig.Digits = big.Digits[0:digitsOfSmaller]
	necessaryBig.Older = small.Older
	if Compare(necessaryBig, small) == 1 {
		necessaryBig.Digits = append(necessaryBig.Digits, big.Digits[digitsOfSmaller])
		necessaryBig.Older++
		k--
	}
	// умножаем small пока он не станет больше big
	smallMultiplied := CopyN(small)
	for multiplier = 1; Compare(smallMultiplied, necessaryBig) != 2; multiplier += 1 {
		smallMultiplied = Addition(smallMultiplied, small)
	}
	res.MakeN([]uint8{multiplier - 1})
	res = MultiplicationBy10k(res, int(k))
	return res
}

// Addition Семёнов
// Сложение двух наутральных чисел
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

// Subtraction Комаровский
// Вычитание из первого большего натурального числа второго меньшего или равного
func Subtraction(x1, x2 Natural) Natural {
	var a, b, res Natural
	var i, j, k int64
	var mass []uint8
	// опрределяем большее число
	if Compare(x1, x2) == 0 { // если равны сразу возвращаем 0
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
	i = int64(a.Older) //самый младший разряд большего числа
	j = int64(b.Older) //самый младший разряд меньшего числа
	for j > -1 {       // пока не дойдем до старшего числа
		switch {
		case int8(a.Digits[i]-b.Digits[j]) < 0: // если разряд большего числа меньше другого
			if int8(a.Digits[i-1]) > 0 { // если следующий разряд не равен 0 то просто вычитаем из него 1
				a.Digits[i-1] = a.Digits[i-1] - 1

			} else {
				// для случая если идет много нулей подряд их заменяем на 9 ,а из ненулевого вычитаем 1
				k = i - 1
				for a.Digits[k] <= 0 {

					a.Digits[k] = 9
					k -= 1
				}
				a.Digits[k] = a.Digits[k] - 1

			}
			mass = append([]uint8{a.Digits[i] - b.Digits[j] + 10}, mass...)
			// добавляем в начало массив значение разряда
		default:
			mass = append([]uint8{a.Digits[i] - b.Digits[j]}, mass...)
		}
		i -= 1
		j -= 1
	}
	for i > -1 { // если не прошлись по всем разрядам большего числа
		mass = append([]uint8{a.Digits[i]}, mass...)
		i -= 1
	}
	res.MakeN(mass)
	return res
}

// DifferenceOfNaturals Комаровский
// Вычитание из натурального другого натурального, умноженного на цифру для случая с неотрицательным результатом
func DifferenceOfNaturals(x1, x2 Natural, k uint8) Natural {
	var a, b, res Natural
	var mass []uint8
	// опрределяем большее число
	if Compare(x1, x2) == 0 { // если равны сразу возвращаем 0
		if k == 1 { // если они раны и число,на которое необходимо умножить рано одному, сразу возвращаем 0
			res.MakeN(mass)
		}
		return res // иначе возвращается пустой
	}
	a = CopyN(x1)
	b = CopyN(x2)
	b = MultiplicationNaturalNumber(b, k) //умножаем меньшее натуральное число на  заданное
	if Compare(b, a) != 2 {               // если при умножение меньшего на цифру оно не становится больше другого,
		// то вычитаем,если нет ,то возвращается пустой
		res = Subtraction(a, b)
	}
	return res
}

func Multiplication(x Natural, y Natural) Natural {
	var otv Natural //структура для ответа
	var i uint32
	var pow int
	var masSum []Natural //массив структур для сложения
	var change Natural
	otv.Digits = append(otv.Digits, 0) // даем структуре otv нолик чтобы в ней было хоть что-то
	pow = 0
	if (x.Digits[0] == 0 && x.Older == 0) || (y.Digits[0] == 0 && y.Older == 0) { // если хоть один множитель 0 - возвращаем 0
		return otv //otv у нас уже нолик
	}
	if x.Older == 0 && y.Older == 0 { //если оба однозначные
		otv.Digits = append(otv.Digits, (x.Digits[0]*y.Digits[0])/10) //запишем старший разряд
		otv.Digits = append(otv.Digits, (x.Digits[0]*y.Digits[0])%10) // и младший
		otv.MakeN(otv.Digits)                                         //избавимся от нулей
		return otv
	}
	if y.Older == 0 { //если только второй множитель однозначный поменяем местами
		change = y
		y = x
		x = change
	}

	for i = y.Older; i != 0; i-- {
		k := MultiplicationNaturalNumber(x, y.Digits[i]) //умножаем 1 число на каждую цифру второго
		e := MultiplicationBy10k(k, pow)                 //умножаем произведение на степень 10 т.к. сдвигаем влево (ну вы поняли да)
		pow += 1                                         //увеличиваем степень 10
		masSum = append(masSum, e)                       //заполняем массив структур
	}

	k := MultiplicationNaturalNumber(x, y.Digits[i]) //добавляем последнее произведение
	e := MultiplicationBy10k(k, pow)
	masSum = append(masSum, e)
	otv = Addition(otv, e) //сразу прибавляем к ответу последнее произвденеие
	for i = 0; i < uint32(len(masSum)-2); i++ {
		otv = Addition(otv, masSum[i]) //складываем все произведения в массиве
	}
	otv = Addition(otv, masSum[i]) // прибавляем последнее оставшееся
	return otv
}
