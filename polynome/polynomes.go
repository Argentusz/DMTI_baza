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

//MultiplicationRational Семёнов Максим
//Умножение полинома на рациональное число
func MultiplicationRational(a Polynomial, b rational.Rational) Polynomial {
	var i uint32

	for i = 0; i < a.Older; i++ { //прогоняем каждый член полинома(так можно говорить?) отдельно
		a.Coeff[i] = rational.Multiplication(a.Coeff[i], b) // умножаем опред член на рациональное число
	}

	return a
}

//Compare Турбина
//Сравнение полиномов:: 2 - если первое больше второго, 0, если равно, 1 иначе.
func Compare(num1, num2 Polynomial) int {
	var i uint32
	switch {
	case num1.Older > num2.Older: //если старшая степень больше, то полином больше
		return 2
	case num2.Older > num2.Older:
		return 1
	default:
		//сравниваем коэффициенты, если соответствующий коэффициент больше, то одно число больше другого
		for i = 0; i < num1.Older+1; i++ {
			switch {
			case rational.Compare(num1.Coeff[i], num2.Coeff[i]) == 2:
				return 2
			case rational.Compare(num1.Coeff[i], num2.Coeff[i]) == 1:
				return 1
			}
		}
	}

	return 0
}
