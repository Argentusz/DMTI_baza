package polynome

import (
	"DMTI_baza/natural"
	"DMTI_baza/rational"
	"DMTI_baza/whole"
	"fmt"
	"strconv"
)

// Polynomial Максим Тростин
// Структура Полиномов
type Polynomial struct {
	Older uint32
	Coeff []rational.Rational
}

// MakeP Тростин Максим/Голубев Михаил
// Метод создания полинома
func (p *Polynomial) MakeP(coeffs []rational.Rational) {
	for _, v := range coeffs {
		if whole.Positivity(v.Nominator) != 0 { //Проверка на 0 старшего коэфф
			break
		} else {
			coeffs = coeffs[1:]
		}
	}
	p.Coeff = coeffs
	p.Older = uint32(len(coeffs))
	if p.Older == 0 {
		p.Coeff = append(p.Coeff, rational.Zero())
	} else {
		p.Older -= 1
	}
}

// CopyP Максим Тростин
// Функция копирования полинома
func CopyP(p Polynomial) Polynomial {
	var copyP Polynomial
	var i uint32
	copyP.Older = p.Older
	for i = 0; i <= p.Older; i++ {
		copyP.Coeff = append(copyP.Coeff, rational.CopyR(p.Coeff[i]))
	}
	return copyP
}

//Голубев Михаил - AdditionP сложение многочленов
func AdditionP(p1Old Polynomial, p2Old Polynomial) Polynomial {
	var result Polynomial
	var coeffsRes []rational.Rational
	p1, p2 := CopyP(p1Old), CopyP(p2Old)
	if p1.Older == p2.Older { //коэффициенты равны - просто складываем коэффициенты попарно
		for i, v := range p1.Coeff {
			coeffsRes = append(coeffsRes, rational.Addition(v, p2.Coeff[i]))
		}
		result.MakeP(coeffsRes)
		return result
	}
	if p1.Older > p2.Older { //иначе считаем разницу, записываем нв новый полином все коэффициенты большего а после складываем начиная с разницы
		difference := int(p1.Older - p2.Older)
		for i := 0; i < difference; i++ {
			coeffsRes = append(coeffsRes, p1.Coeff[i])
		}
		for i := difference; i <= int(p1.Older); i++ {
			coeffsRes = append(coeffsRes, rational.Addition(p1.Coeff[i], p2.Coeff[i-difference]))
		}
		result.MakeP(coeffsRes)
	}

	if p1.Older < p2.Older {
		difference := int(p2.Older - p1.Older)
		for i := 0; i < difference; i++ {
			coeffsRes = append(coeffsRes, p2.Coeff[i])
		}
		for i := difference; i <= int(p2.Older); i++ {
			coeffsRes = append(coeffsRes, rational.Addition(p2.Coeff[i], p1.Coeff[i-difference]))
		}
		result.MakeP(coeffsRes)
	}
	return result
}

//Голубев Михаил - SubstractionPol вычитание многочленов
func SubstractionP(fromOld Polynomial, whatOld Polynomial) Polynomial {
	var result Polynomial
	var coeffsRes []rational.Rational
	from, what := CopyP(fromOld), CopyP(whatOld)
	if from.Older == what.Older { //если степени равны - просто вычитаем коэффициенты
		for i, v := range from.Coeff {
			coeffsRes = append(coeffsRes, rational.Subtraction(v, what.Coeff[i]))
		}
		result.MakeP(coeffsRes)
		return result
	}
	if from.Older > what.Older { //если коэффициенты не равны - считаем разницу и в новый полином записываем все коэффициенты большего до разницы
		difference := int(from.Older - what.Older)
		for i := 0; i <= difference; i++ {
			coeffsRes = append(coeffsRes, from.Coeff[i])
		}
		for i := difference; i < int(from.Older); i++ { //и вычитаем коэффициенты после разницы
			coeffsRes = append(coeffsRes, rational.Subtraction(from.Coeff[i], what.Coeff[i-difference]))
		}
		result.MakeP(coeffsRes)
	}
	if from.Older < what.Older {
		difference := int(what.Older - from.Older)
		for i := 0; i < difference; i++ {
			if what.Coeff[i].Nominator.Negative != true {
				what.Coeff[i].Nominator.Negative = true
			} else {
				what.Coeff[i].Nominator.Negative = false
			}
			coeffsRes = append(coeffsRes, what.Coeff[i])
		}
		for i := difference; i <= int(what.Older); i++ {
			coeffsRes = append(coeffsRes, rational.Subtraction(what.Coeff[i], from.Coeff[i-difference]))
		}
		result.MakeP(coeffsRes)
	}
	return result
}

//Голубев Михаил - MultiplicationXpowerK функция умножения многочлена на x^k
func MultiplicationXpowerK(polynome Polynomial, k int) Polynomial {
	nullCoeff := rational.Rational{Nominator: whole.Whole{Num: natural.Natural{Digits: []uint8{0}, Older: 0},
		Negative: false}, Denominator: natural.Natural{Digits: []uint8{1}, Older: 0}} //создаем коэффициент вида 0/1
	p := CopyP(polynome)
	coeffs := p.Coeff
	var result Polynomial
	for i := 0; i < k; i++ {
		coeffs = append(coeffs, nullCoeff) //обнуляем все коэффициенты меньше x^k
	}
	result.MakeP(coeffs)
	return result
}

// Голубев Михаил OlderPoly Степень многочлена
func OlderPoly(p Polynomial) uint32 {
	return p.Older
}

// Голубев Михаил OlderCoeffPoly Старший коэффициент многочлена
func OlderCoeffPoly(p Polynomial) rational.Rational {
	return p.Coeff[0]
}

// Derivative Максим Тростин
// Производная многочлена
func Derivative(p0 Polynomial) Polynomial {
	var i uint32
	p := CopyP(p0)
	// Константа просто уходит
	p.Coeff[p.Older] = rational.Zero()
	for i = p.Older; i != 0; i-- {
		// Каждый коэффициент - это произведение степени старшего на коэффициент старшего
		p.Coeff[i] = rational.Multiplication(p.Coeff[i-1], rational.IntToRational(int64(p.Older-i+1), 1))
	}
	// Убираем оставшуюся старшую степень
	p.Older--
	var _ rational.Rational
	_, p.Coeff = p.Coeff[0], p.Coeff[1:]
	return p
}

//Compare Турбина
//Сравнение полиномов: 0, если равно, 1 иначе.
func Compare(num1, num2 Polynomial) int {
	var i uint32
	switch {
	case num1.Older > num2.Older: //если старшая степень больше, то полином больше
		return 1
	case num2.Older > num1.Older:
		return 1
	default:
		//сравниваем коэффициенты, если соответствующий коэффициент больше, то одно число больше другого
		for i = 0; i < num1.Older+1; i++ {
			if rational.Compare(num1.Coeff[i], num2.Coeff[i]) != 0 {
				return 1
			}
		}
	}

	return 0
}

//MultiplicationRational Семёнов Максим
//Умножение полинома на рациональное число
func MultiplicationRational(a Polynomial, b rational.Rational) Polynomial {
	var i uint32

	for i = 0; i <= a.Older; i++ { //прогоняем каждый член полинома(так можно говорить?) отдельно
		a.Coeff[i] = rational.Multiplication(a.Coeff[i], b) // умножаем опред член на рациональное число
	}

	return a
}

// ToStringPol Максим Тростин.
// Возвращает полином в строковом виде
func (p *Polynomial) ToStringPol() string {
	var str string
	var i uint32
	if p.Older == 0 {
		return fmt.Sprint(p.Coeff[0])
	} else {
		for i = 0; i < p.Older; i++ {
			if whole.Positivity(p.Coeff[i].Nominator) == 2 {
				str += "+" + rational.ToStringR(p.Coeff[i]) + "x" + "^" + strconv.FormatUint(uint64(p.Older-i), 10)
			} else if whole.Positivity(p.Coeff[i].Nominator) == 1 {
				str += rational.ToStringR(p.Coeff[i]) + "x" + "^" + strconv.FormatUint(uint64(p.Older-i), 10)
			}

		}
	}
	if str[0] == '+' {
		str = str[1:]
	}
	return str
}

// Пименов
// НОД многочленов
func GreatestCommonDivisor(aOriginal, bOriginal Polynomial) Polynomial {
	var a, b Polynomial                        // Копии входных
	var GCD Polynomial                         // Наибольший общий делитель (результат)
	var nominator, denominator natural.Natural // Множитель и делитель для полученного многочлена
	var multiplier rational.Rational           // Множитель для полученного многочлена
	// Дело в том, что НОД для многочлена неоднозначен и может умножаться на C
	a = CopyP(aOriginal)
	b = CopyP(bOriginal)
	// Применяем алгоритм Евклида
	for !((a.Older == 0 && rational.Compare(a.Coeff[0], rational.Zero()) == 0) ||
		(b.Older == 0 && rational.Compare(b.Coeff[0], rational.Zero()) == 0)) {
		if a.Older >= b.Older {
			a = RemainderFromDivision(a, b)
		} else {
			b = RemainderFromDivision(b, a)
		}
	}
	GCD = AdditionP(a, b)

	// Будем делить на НОД числителей и умножать на НОК знаменателей
	denominator, nominator = GreatestCommonDivisorAndLeastCommonMultipleOfPolynomial(GCD)
	multiplier.MakeR(whole.FromNaturalsToWhole(nominator), denominator)
	if GCD.Coeff[0].Nominator.Negative == true { // Старший коэфф. делаем с плюсом для красоты
		multiplier.Nominator.Negative = !multiplier.Nominator.Negative
	}
	GCD = MultiplicationRational(GCD, multiplier) // Умножаем на наш специальный множитель
	return GCD
}

// QuotientOfDivision Комаровский
// Частное от деления полиномов
func QuotientOfDivision(x1, x2 Polynomial) Polynomial {
	var coef rational.Rational
	var diff uint32
	var x, y, mid, res Polynomial
	var midcoef []rational.Rational
	x = CopyP(x1)
	if x2.Older == 0 { //если нужно поделить просто на число
		for i := 0; i <= int(x.Older); i++ {
			x.Coeff[i] = rational.Division(x.Coeff[i], x2.Coeff[0])
		}
		return x
	}
	y = CopyP(x2)
	for x.Older >= y.Older { // пока стрепень первого больше или равна второго
		diff = x.Older - y.Older                         // разница их степеней
		coef = rational.Division(x.Coeff[0], y.Coeff[0]) //деление старшего коэфицента первого и второго для определения коэфицента в частном
		mid = MultiplicationXpowerK(y, int(diff))        //домножаем на степень
		mid = MultiplicationRational(mid, coef)
		x = SubstractionP(x, mid) // вычитаем, тем самым убирается старшая степень
		midcoef = append(midcoef, coef)
	}
	res.MakeP(midcoef)
	return res
}

//Грунская Умножение полиномов

func MultiplicationPol(xOld, yOld Polynomial) Polynomial {
	var x, y, otv Polynomial
	var i uint32
	var SumMas []Polynomial
	x = CopyP(xOld) //делаем копии на всякий случай,чтобы не было казусов
	y = CopyP(yOld)
	otv.MakeP([]rational.Rational{rational.Zero()})
	for i = 0; i < y.Older+1; i++ {
		k := MultiplicationXpowerK(x, int(i))              //умножаем на х^л
		e := MultiplicationRational(k, y.Coeff[y.Older-i]) //умножаем на коэффициент
		SumMas = append(SumMas, e)                         //заносим в массив для последующего сложения
	}
	for i = 0; int(i) < len(SumMas); i++ {
		otv = AdditionP(otv, SumMas[i]) // прибавляем
	}
	return otv
}

// Остаток от деления полиномов
func RemainderFromDivision(x1, x2 Polynomial) Polynomial {
	var x, y, mid, res Polynomial
	x, y = CopyP(x1), CopyP(x2)
	mid = QuotientOfDivision(x1, x2)                  // частное от деления
	res = SubstractionP(x, MultiplicationPol(mid, y)) // разность Делимого и умножения делителя на частного
	return res
}

//Морозов никита
//Вынесение НОД числителей и НОК знаменателей
//Как работает функция:
//Находим НОД числителей
//Находим НОК знаменателей
//Возвращаем результат как два натуральных числа: НОД и НОК
func GreatestCommonDivisorAndLeastCommonMultipleOfPolynomial(polynom Polynomial) (natural.Natural, natural.Natural) {
	//Объявляем переменные НОД (GCD) и НОК (LCM)
	var GCD []natural.Natural
	var LCM []natural.Natural

	//Клонироуем все элементы числители в массив НОД
	for i := 0; i < len(polynom.Coeff); i++ {
		//Исключаем нулевые коэффициенты
		if natural.CheckNull(whole.Absolute(polynom.Coeff[i].Nominator)) == false {
			GCD = append(GCD, whole.Absolute(polynom.Coeff[i].Nominator))
		}
	}

	//Клонируем все элементы знаменателя в массив НОК
	for i := 0; i < len(polynom.Coeff); i++ {
		LCM = append(LCM, polynom.Coeff[i].Denominator)
	}

	//Находим НОД
	//Находим НОД от числителей под номерами 0 - 1, 1 - 2, 2 - 3... (n-1) - n
	//Записываем в ячейки массива
	//В конец последняя ячейка не изменится — её удаляем
	//Повторяем пока массив не будет содержать только один элемент
	//Пример:
	//Числители 3 6 9
	//Находим НОД от 3 и 6: 3
	//Находим НОД от 6 и 9: 3
	//Записываем в массив: 3 меняется на 3, 6 меняется на 3, девять остаётся
	//Измененённый массив: 3 3 9
	//Удаляем 9
	//Получается 3 3
	//Находим НОД снова, получается массив 3 3
	//Удаляем последнюю 3
	//Ответ 3
	//Итоговая пирамидка:
	//3 6 9
	//3 3 [9]
	//3 3
	//3 [3]
	//3
	for len(GCD) != 1 {
		for i := 0; i < len(GCD)-1; i++ {
			GCD[i] = natural.GreatestCommonDivisor(GCD[i], GCD[i+1])
		}
		GCD = GCD[:len(GCD)-1]
	}

	//Находим НОК по аналогии с НОД
	for len(LCM) != 1 {
		for i := 0; i < len(LCM)-1; i++ {
			LCM[i] = natural.LeastCommonMultiple(LCM[i], LCM[i+1])
		}
		LCM = LCM[:len(LCM)-1]
	}

	return GCD[0], LCM[0]
}
