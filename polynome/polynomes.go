package polynome

import (
	"DMTI_baza/rational"
	"DMTI_baza/whole"
	"fmt"
	"math"
	"strconv"
)

//----------НЕ ЗАБЫТЬ ПОМЕНЯТЬ FLOAT----------//
//Структура и методы - от Голубева Михаила

type Polynomial struct {
	Older uint32
	Coeff []rational.Rational
}

// Old структура полиномов //
//каждый полином имеет старшую степень и коэффициенты при степенях
type Old struct {
	Older  int
	Coeffs []float64
}

//------------------МЕТОДЫ------------------//

// MakePolOld Конструктор полинома - метод
func (p *Old) MakePolOld(coeffs []float64) {
	for _, v := range coeffs {
		if math.Round(v*10000) != 0.0 { //Проверка на 0 старшего коэфф
			break
		} else {
			coeffs = coeffs[1:]
		}
	}
	p.Coeffs = coeffs
	p.Older = len(coeffs) - 1
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

// ToStringPolOld Приведение полинома к строковому виду "ax^n+bx^(n-1)+...+nx^0" - метод
func (p *Old) ToStringPolOld() string {
	var str string
	x := "x"
	if p.Older == 0 {
		return fmt.Sprint(p.Coeffs[0])
	} else {
		for i := 0; i < len(p.Coeffs); i++ {
			if p.Coeffs[i] > 0 {
				str += "+" + fmt.Sprint(p.Coeffs[i]) + x + "^" + strconv.Itoa(p.Older-i)
			} else if p.Coeffs[i] == 0 {

			} else if p.Coeffs[i] < 0 {
				str += fmt.Sprint(p.Coeffs[i]) + x + "^" + strconv.Itoa(p.Older-i)
			}

		}
	}
	return str
}

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
				str += "-" + rational.ToStringR(p.Coeff[i]) + "x" + "^" + strconv.FormatUint(uint64(p.Older-i), 10)
			}

		}
	}
	if str[0] == '+' {
		str = str[1:]
	}
	return str
}

//------------------ФУНКЦИИ------------------//

//Функция - от Голубева Михаила

// AdditionPolOld Сложение многочленов
func AdditionPolOld(p1 Old, p2 Old) Old {
	var ans Old
	var arr []float64
	var lenDiffs int
	if p1.Older == p2.Older { //степени равны - просто складываем коэффициенты
		ans.Older = p1.Older
		for index := range p1.Coeffs {
			val := p2.Coeffs[index] + p1.Coeffs[index]
			arr = append(arr, val)
		}
		ans.MakePolOld(arr)
	} else if p1.Older > p2.Older { //степень первого больше
		ans.Older = p1.Older
		lenDiffs = len(p1.Coeffs) - len(p2.Coeffs)
		for i := 0; i < lenDiffs; i++ { //сначала записываем коэффициенты до разницы степеней
			val := p1.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range p2.Coeffs { //записываем сумму остальных
			val := v + p1.Coeffs[i+lenDiffs]
			arr = append(arr, val)
		}
		ans.MakePolOld(arr)
	} else if p1.Older < p2.Older { //аналогично, но наоборот
		lenDiffs = len(p2.Coeffs) - len(p1.Coeffs)
		ans.Older = p2.Older
		for i := 0; i < lenDiffs; i++ {
			val := p2.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range p1.Coeffs {
			val := v + p2.Coeffs[i+lenDiffs]
			arr = append(arr, val)
		}
		ans.MakePolOld(arr)
	}
	return ans
}

//Функция - от Голубева Михаила

// SubtractionPolOld Вычитание многочленов
func SubtractionPolOld(from Old, what Old) Old {
	var ans Old
	var arr []float64
	var lenDiffs int
	if from.Older == what.Older { //если они одной степени - просто вычитаем коэффициенты
		ans.Older = from.Older
		for i, v := range from.Coeffs {
			val := v - what.Coeffs[i]
			arr = append(arr, val)
		}
		ans.MakePolOld(arr)
	} else if from.Older > what.Older { //степень того из которого вычитаем больше степени вычитаемого
		lenDiffs = len(from.Coeffs) - len(what.Coeffs)
		ans.Older = from.Older
		for i := 0; i < lenDiffs; i++ { //сначала записываем все коэффициенты, идущие до разницы степеней
			val := from.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range what.Coeffs {
			val := from.Coeffs[i+lenDiffs] - v //записываем разность остальных коэффов
			arr = append(arr, val)
		}
		ans.MakePolOld(arr)
	} else if from.Older < what.Older { //наоборот - степень второго больше
		lenDiffs = len(what.Coeffs) - len(from.Coeffs)
		ans.Older = what.Older
		for i := 0; i < lenDiffs; i++ { //записываем противоположные по знаку коэффициенты до разницы степеней
			val := 0 - what.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range from.Coeffs { //записываем разность остальных коэфф.
			val := v - what.Coeffs[i+lenDiffs]
			arr = append(arr, val)
		}
		ans.MakePolOld(arr)
	}
	return ans
}

//Функция - от Голубева Михаила

// NumberMultiplyPolOld  Умножение многочлена на рациональное число
func NumberMultiplyPolOld(p Old, a float64) Old {
	for index := range p.Coeffs {
		p.Coeffs[index] *= a
	}
	return p
}

//Функция - от Голубева Михаила

// VarMultiplyPolOld Умножение многочлена на x^k
func VarMultiplyPolOld(p Old, k int) Old {
	p.Older += k
	for i := 0; i < p.Older; i++ {
		p.Coeffs = append(p.Coeffs, 0) //Добавляем нули как коэффициенты младших степеней
	}
	return p
}

//Функция - от Голубева Михаила

// OlderPoly Степень многочлена
func OlderPoly(p Polynomial) uint32 {
	return p.Older
}

//Функция - от Голубева Михаила

// OlderCoeffPoly Старший коэффициент многочлена
func OlderCoeffPoly(p Polynomial) rational.Rational {
	return p.Coeff[0]
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
		GCD = append(GCD, whole.Absolute(polynom.Coeff[i].Nominator))
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
