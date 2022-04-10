package polynome

import (
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

	for i = 0; i < a.Older; i++ { //прогоняем каждый член полинома(так можно говорить?) отдельно
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
