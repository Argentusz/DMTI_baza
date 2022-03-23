package polynome

import (
	"fmt"
	"math"
	"strconv"
)

//Структура и методы - от Голубева Михаила

//структура полиномов //
//каждый полином имеет старшую степень и коэффициенты при степенях
type Polynome struct {
	Older  int
	Coeffs []float64
}

//------------------МЕТОДЫ------------------//

//Конструктор полинома - метод
////Замечания по неймингу приветствуются
func (p *Polynome) MakePol(coeffs []float64) {
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

//Приведение полинома к строковому виду "ax^n+bx^(n-1)+...+nx^0" - метод
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

//------------------ФУНКЦИИ------------------//

//Функция - от Голубева Михаила

//Сложение многочленов
func AdditionPol(p1 Polynome, p2 Polynome) Polynome {
	var ans Polynome
	var arr []float64
	var lenDiffs int
	if p1.Older == p2.Older { //степени равны - просто складываем коэффициенты
		ans.Older = p1.Older
		for index := range p1.Coeffs {
			val := p2.Coeffs[index] + p1.Coeffs[index]
			arr = append(arr, val)
		}
		ans.MakePol(arr)
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
		ans.MakePol(arr)
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
		ans.MakePol(arr)
	}
	return ans
}

//Функция - от Голубева Михаила

//Вычитание многочленов
func SubtractionPol(from Polynome, what Polynome) Polynome {
	var ans Polynome
	var arr []float64
	var lenDiffs int
	if from.Older == what.Older { //если они одной степени - просто вычитаем коэффициенты
		ans.Older = from.Older
		for i, v := range from.Coeffs {
			val := v - what.Coeffs[i]
			arr = append(arr, val)
		}
		ans.MakePol(arr)
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
		ans.MakePol(arr)
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
		ans.MakePol(arr)
	}
	return ans
}

//Функция - от Голубева Михаила

//Умножение многочлена на рациональное число
func NumberMultiplyPol(p Polynome, a float64) Polynome {
	for index := range p.Coeffs {
		p.Coeffs[index] *= a
	}
	return p
}

//Функция - от Голубева Михаила

//Умножение многочлена на x^k
func VarMultiplyPol(p Polynome, k int) Polynome {
	p.Older += k
	for i := 0; i < p.Older; i++ {
		p.Coeffs = append(p.Coeffs, 0) //Добавляем нули как коэффициенты младших степеней
	}
	return p
}

//Функция - от Голубева Михаила

//Степень многочлена
func OlderPoly(p Polynome) int {
	return p.Older
}

//Функция - от Голубева Михаила

//Старший коэффициент многочлена
func OlderCoeffPoly(p Polynome) float64 {
	return p.Coeffs[0]
}
