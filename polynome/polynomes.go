package polynome

import (
	"fmt"
	"strconv"
)

type Polynome struct {
	Older  int
	Coeffs []float64
}

func (p *Polynome) MakePol(coeffs []float64) {
	p.Coeffs = coeffs
	p.Older = len(coeffs) - 1
}

func (p *Polynome) ToString() string {
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

//Сложение многочленов
func AdditionPol(p1 Polynome, p2 Polynome) Polynome {
	var ans Polynome
	var arr []float64
	var lenDiffs int
	if p1.Older == p2.Older {
		ans.Older = p1.Older
		for i, v := range p1.Coeffs {
			val := p2.Coeffs[i] + v
			arr = append(arr, val)
		}
		ans.MakePol(arr)
	} else if p1.Older > p2.Older {
		ans.Older = p1.Older
		lenDiffs = len(p1.Coeffs) - len(p2.Coeffs)
		for i := 0; i < lenDiffs; i++ {
			val := p1.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range p2.Coeffs {
			val := v + p1.Coeffs[i+lenDiffs]
			arr = append(arr, val)
		}
		ans.MakePol(arr)
	} else if p1.Older < p2.Older {
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
		ans.MakePol(arr)
	}
	return ans
}

//Вычитание многочленов
func SubtractionPol(from Polynome, what Polynome) Polynome {
	var ans Polynome
	var arr []float64
	var lenDiffs int
	if from.Older == what.Older {
		ans.Older = from.Older
		for i, v := range from.Coeffs {
			val := v - what.Coeffs[i]
			arr = append(arr, val)
		}
		ans.MakePol(arr)
	} else if from.Older > what.Older {
		lenDiffs = len(from.Coeffs) - len(what.Coeffs)
		ans.Older = from.Older
		for i := 0; i < lenDiffs; i++ {
			val := from.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range what.Coeffs {
			val := from.Coeffs[i+lenDiffs] - v
			arr = append(arr, val)
		}
		ans.MakePol(arr)
	} else if from.Older < what.Older {
		lenDiffs = len(what.Coeffs) - len(from.Coeffs)
		ans.Older = what.Older
		for i := 0; i < lenDiffs; i++ {
			val := 0 - what.Coeffs[i]
			arr = append(arr, val)
		}
		for i, v := range from.Coeffs {
			val := v - what.Coeffs[i+lenDiffs]
			arr = append(arr, val)
		}
		ans.MakePol(arr)
	}
	return ans
}

//Умножение многочлена на рациональное число
func NumberMultiplyPol(p Polynome, a float64) Polynome {
	for index := range p.Coeffs {
		p.Coeffs[index] *= a
	}
	return p
}

//Умножение многочлена на x^k
func VarMultiplyPol(p Polynome, k int) Polynome {
	p.Older += k
	return p
}

//Старший коэффициент многочлена
func OlderPoly(p Polynome) int {
	return p.Older
}

//Степень многочлена
func OlderCoeffPoly(p Polynome) float64 {
	return p.Coeffs[0]
}
