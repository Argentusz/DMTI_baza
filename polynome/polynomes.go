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
