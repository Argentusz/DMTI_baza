package main

//Исполняемоый пакет для работы с консольными флагами
//В целом, на него смотреть пока не нужно, это я сам добью, потом посмотрите и оцените(!)
import (
	"DMTI_baza/natural"
	"DMTI_baza/whole"
	"fmt"
	"strconv"
	"strings"
)

//Модуль от Голубева Михаила
func main() {
	var moduleNumber int
	var functionNumber int
	fmt.Println("введите номер модуля")
	fmt.Scan(&moduleNumber)
	fmt.Println("введите номер функции")
	fmt.Scan(&functionNumber)
	switch moduleNumber {
	case 1:
		naturalsGui(functionNumber)
	case 2:

	}
}

func naturalsGui(f int) {
	switch f {
	case 1: //сравнение
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.Compare(n1, n2))
	case 2: //проверка на 0
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println(natural.CheckNull(n1))
	case 3: //доавление 1
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println(natural.ToStringN(natural.Addition1(n1)))
	case 4: //сложение
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.ToStringN(natural.Addition(n1, n2)))
	case 5: //Вычитание из первого большего натурального числа второго меньшего или равного
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		res := natural.Compare(n1, n2)
		if res == 1 {
			fmt.Println("wrong")
			return
		}
		fmt.Println(natural.ToStringN(natural.Subtraction(n1, n2)))
	case 6: //умножение на цифру
		fmt.Println("first")
		var s string
		var k uint8
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&k)
		fmt.Println(natural.ToStringN(natural.MultiplicationNaturalNumber(n1, k)))
	case 7: //Умножение натурального числа на 10^k
		fmt.Println("first")
		var s string
		var k int
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&k)
		fmt.Println(natural.ToStringN(natural.MultiplicationBy10k(n1, k)))
	case 8: //Умножение натуральных чисел
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.Multiplication(n1, n2))
	case 9: //Вычитание из натурального другого натурального, умноженного на цифру для случая с неотрицательным результатом
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		fmt.Println("k=")
		var k uint8
		fmt.Scan(&k)
		n2 := inputNatural(s)
		fmt.Println(natural.DifferenceOfNaturals(n1, n2, k))
	case 10: //Вычисление первой цифры деления большего натурального на меньшее, домноженное на 10^k,где k - номер позиции этой цифры (номер считается с нуля)
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.DivideOneIteration(n1, n2))
	case 11: //Частное от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля)
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.IntegerFromDivision(n1, n2))
	case 12: //Остаток от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля)
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.RemainderFromDivision(n1, n2))
	case 13: //нод
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.GreatestCommonDivisor(n1, n2))
	case 14: //НОК
		fmt.Println("first")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("second")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.LeastCommonMultiple(n1, n2))
	}
}

func inputNatural(s string) natural.Natural {
	nat := strings.Split(s, "")
	var natur []uint8
	for _, v := range nat {
		n, _ := strconv.Atoi(v)
		natur = append(natur, uint8(n))
	}
	var n natural.Natural
	n.MakeN(natur)
	return n
}

func wholeGui(f int) {
	var s string
	switch f {
	case 1: //Абсолютная величина числа, результат - натуральное
		fmt.Println("whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(natural.ToStringN(whole.Absolute(w1)))
	case 2: //Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)
		fmt.Println("whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(whole.Positivity(w1))
	case 3: //Умножение целого на (-1)
		fmt.Println("whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.MultiplicationByNegativeOne(w1)))
	case 4: //Преобразование натурального в целое
		fmt.Println("first")
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println(whole.ToStringW(whole.FromNaturalsToWhole(n1)))
	case 5: //Преобразование целого неотрицательного в натуральное
		fmt.Println("first")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(natural.ToStringN(whole.FromWholeToNaturals(w1)))
	case 6: //Сложение целых чисел
		fmt.Println("1st whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("2nd whole-")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.Addition(w1, w2)))
	case 7: //вычитание целых чисел
		fmt.Println("from whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("to whole-")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.Subtraction(w1, w2)))
	case 8: //умножение целых чисел
		fmt.Println("1st whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("2nd whole-")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.Multiplication(w1, w2)))
	case 9: //Частное от деления целого на целое (делитель отличен от нуля)
		fmt.Println("1st whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("2nd whole-")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.WholeFromDivision(w1, w2)))
	case 10: //Остаток от деления целого на целое(делитель отличен от нуля)
		fmt.Println("1st whole-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("2nd whole-")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.RemainderFromDivision(w1, w2)))
	}
}

func inputWhole(s string) whole.Whole {
	var negative = false
	if s[0] == '-' {
		negative = true
	}
	s = s[1:]
	nums := inputNatural(s).Digits
	var w whole.Whole
	w.MakeW(negative, nums)
	return w
}
