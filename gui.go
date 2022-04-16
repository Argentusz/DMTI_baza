package main

// GUI Михаил Голубев & Максим Тростин
import (
	"DMTI_baza/natural"
	"DMTI_baza/polynome"
	"DMTI_baza/rational"
	"DMTI_baza/whole"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func main() {
	// Functions names in modules
	funcsN := []string{
		"Сравнение натуральных чисел: 2 - если первое больше второго, 0, если равно, 1 иначе.",
		"Проверка на ноль: если число не равно нулю, то «да» иначе «нет»",
		"Добавление 1 к натуральному числу",
		"Сложение натуральных чисел",
		"Вычитание из первого большего натурального числа второго меньшего или равного",
		"Умножение натурального числа на цифру",
		"Умножение натурального числа на 10^k",
		"Умножение натуральных чисел",
		"Вычитание из натурального другого натурального, умноженного на цифру для случая с неотрицательным результатом",
		"Вычисление первой цифры деления большего натурального на меньшее, домноженное на 10^k,где k - номер позиции этой цифры",
		"Частное от деления большего натурального числа на меньшее или равное натуральное с остатком",
		"Остаток от деления большего натурального числа на меньшее или равное натуральное с остатком",
		"НОД натуральных чисел",
		"НОК натуральных чисел",
	}
	funcsWh := []string{
		"Абсолютная величина числа, результат - натуральное",
		"Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)",
		"Умножение целого на (-1)",
		"Преобразование натурального в целое",
		"Преобразование целого неотрицательного в натуральное",
		"Сложение целых чисел",
		"Вычитание целых чисел",
		"Умножение целых чисел",
		"Частное от деления целого на целое (делитель отличен от нуля)",
		"Остаток от деления целого на целое(делитель отличен от нуля)",
	}

	funcsR := []string{
		"Сокращение дроби",
		"Проверка на целое, если рациональное число является целым, то «да», иначе «нет»",
		"Преобразование целого в дробное",
		"Преобразование дробного в целое (если знаменатель равен 1)",
		"Сложение дробей",
		"Вычитание дробей",
		"Умножение дробей",
		"Деление дробей (делитель отличен от нуля)",
	}

	funcsPol := []string{
		"Сложение многочленов",
		"Вычитание многочленов",
		"Умножение многочлена на рациональное число",
		"Умножение многочлена на x^k",
		"Старший коэффициент многочлена",
		"Степень многочлена",
		"Вынесение из многочлена НОК знаменателей коэффициентов и НОД числителей",
		"Умножение многочленов",
		"Частное от деления многочлена на многочлен при делении с остатком",
		"Остаток от деления многочлена на многочлен при делении с остатком",
		"НОД многочленов",
		"Производная многочлена",
		"Преобразование многочлена — кратные корни в простые",
	}

	// Creating Window
	a := app.New()
	w := a.NewWindow("BAZA Unlimited")
	w.Resize(fyne.NewSize(1000, 500))

	// Modules
	radioMods := widget.NewRadioGroup([]string{"Натуральные", "Целые", "Рациональные", "Полиномы"}, func(s string) {})
	radioMods.SetSelected("Натуральные")

	// Module Pick Menu
	// natural
	funcBoxN := widget.NewRadioGroup(funcsN, func(s string) {})
	funcBoxN.SetSelected(funcsN[0])

	// whole
	funcBoxWh := widget.NewRadioGroup(funcsWh, func(s string) {})
	funcBoxWh.SetSelected(funcsWh[0])

	// rational
	funcBoxR := widget.NewRadioGroup(funcsR, func(s string) {})
	funcBoxR.SetSelected(funcsR[0])

	// polynomial
	funcBoxPol := widget.NewRadioGroup(funcsPol, func(s string) {})
	funcBoxPol.SetSelected(funcsPol[0])

	// Entry Fields
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()
	// Selected Module
	Selected := ""
	// Selected Function
	SelectedFun := "0"
	// Result Text
	result := widget.NewLabel("")
	format := widget.NewLabel("Пример ввода: 4/3x^4+2/1x^2+1/1x")
	entry1.Hide()
	entry2.Hide()
	entry3.Hide()
	format.Hide()
	// Button to Calculate the Result
	btnRes := widget.NewButton("Результат", func() {
		switch SelectedFun {
		// natural
		case "00":
			// Compare
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.Compare(n1, n2)
			result.SetText(strconv.Itoa(res))
		case "01":
			// CheckNull
			n1 := InputNatural(entry1.Text)
			res := natural.CheckNull(n1)
			if res {
				result.SetText("Да")
			} else {
				result.SetText("Нет")
			}
		case "02":
			// Addition1
			n1 := InputNatural(entry1.Text)
			res := natural.Addition1(n1)
			result.SetText(natural.ToStringN(res))
		case "03":
			// Addition
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.Addition(n1, n2)
			result.SetText(natural.ToStringN(res))
		case "04":
			// Subtraction
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.Subtraction(n1, n2)
			result.SetText(natural.ToStringN(res))
		case "05":
			// MultiplicationNaturalNumber
			n1 := InputNatural(entry1.Text)
			if len(entry2.Text) == 1 {
				n2, err := strconv.Atoi(entry2.Text)
				if err != nil {
					result.SetText("Ошибка: x05")
				} else {
					digit := uint8(n2)
					res := natural.MultiplicationNaturalNumber(n1, digit)
					result.SetText(natural.ToStringN(res))
				}
			} else {
				result.SetText("Неправильная длина цифры")
			}
		case "06":
			// MultiplicationBy10k
			n1 := InputNatural(entry1.Text)
			k, err := strconv.Atoi(entry2.Text)
			if err != nil {
				result.SetText("Ошибка: x06")
			} else {
				res := natural.MultiplicationBy10k(n1, k)
				result.SetText(natural.ToStringN(res))
			}
		case "07":
			// Multiplication
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.Multiplication(n1, n2)
			result.SetText(natural.ToStringN(res))
		case "08":
			// DifferenceOfNaturals
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			if len(entry3.Text) == 1 {
				n3, err := strconv.Atoi(entry3.Text)
				if err != nil {
					result.SetText("Ошибка: x08")
				} else {
					digit := uint8(n3)
					res := natural.DifferenceOfNaturals(n1, n2, digit)
					result.SetText(natural.ToStringN(res))
				}
			} else {
				result.SetText("Неправильная длина цифры")
			}
		case "09":
			// DivideOneIteration
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			if natural.CheckNull(n2) {
				result.SetText("Ошибка нулевой делитель: x09")
			} else {
				res := natural.DivideOneIteration(n1, n2)
				result.SetText(natural.ToStringN(res))
			}
		case "0A":
			// IntegerFromDivision
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			if natural.CheckNull(n2) {
				result.SetText("Ошибка нулевой делитель: x0A")
			} else {
				res := natural.IntegerFromDivision(n1, n2)
				result.SetText(natural.ToStringN(res))
			}
		case "0B":
			// RemainderFromDivision
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			if natural.CheckNull(n2) {
				result.SetText("Ошибка нулевой делитель: x09")
			} else {
				res := natural.RemainderFromDivision(n1, n2)
				result.SetText(natural.ToStringN(res))
			}
		case "0C":
			// GreatestCommonDivisor
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.GreatestCommonDivisor(n1, n2)
			result.SetText(natural.ToStringN(res))
		case "0D":
			// LeastCommonMultiple
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.LeastCommonMultiple(n1, n2)
			result.SetText(natural.ToStringN(res))
		// whole
		case "10":
			// Absolute
			w1 := InputWhole(entry1.Text)
			res := whole.Absolute(w1)
			result.SetText(natural.ToStringN(res))
		case "11":
			// Positivity
			w1 := InputWhole(entry1.Text)
			res := whole.Positivity(w1)
			result.SetText(strconv.Itoa(res))
		case "12":
			// MultiplicationByNegativeOne
			w1 := InputWhole(entry1.Text)
			res := whole.MultiplicationByNegativeOne(w1)
			result.SetText(whole.ToStringW(res))
		case "13":
			// FromNaturalsToWhole
			n1 := InputNatural(entry1.Text)
			res := whole.FromNaturalsToWhole(n1)
			result.SetText(whole.ToStringW(res))
		case "14":
			// FromWholeToNaturals
			w1 := InputWhole(entry1.Text)
			res := whole.FromWholeToNaturals(w1)
			result.SetText(natural.ToStringN(res))
		case "15":
			// Addition
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.Addition(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "16":
			// Subtraction
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.Subtraction(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "17":
			// Multiplication
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.Multiplication(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "18":
			// WholeFromDivision
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			if whole.Compare(w2, whole.IntToWhole(0)) == 0 {
				result.SetText("Ошибка нулевого делителя: x18")
			} else {
				res := whole.WholeFromDivision(w1, w2)
				result.SetText(whole.ToStringW(res))
			}
		case "19":
			// RemainderFromDivision
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			if whole.Compare(w2, whole.IntToWhole(0)) == 0 {
				result.SetText("Ошибка нулевого делителя: x18")
			} else {
				res := whole.RemainderFromDivision(w1, w2)
				result.SetText(whole.ToStringW(res))
			}
		case "20":
			// SimplifyingFractions
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x20")
			} else {
				res := rational.SimplifyingFractions(r1)
				result.SetText(rational.ToStringR(res))
			}
		case "21":
			// CheckingForWhole
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x21")
			} else {
				res := rational.CheckingForWhole(r1)
				if res {
					result.SetText("Да")
				} else {
					result.SetText("Нет")
				}
			}
		case "22":
			// WholeToFractional
			w1 := InputWhole(entry1.Text)
			res := rational.WholeToFractional(w1)
			result.SetText(rational.ToStringR(res))
		case "23":
			// FractionalToWhole
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x23")
			} else {
				res := rational.FractionalToWhole(r1)
				result.SetText(whole.ToStringW(res))
			}
		case "24":
			// Addition
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x24")
			} else {
				r2, err := InputRational(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x24")
				} else {
					res := rational.Addition(r1, r2)
					result.SetText(rational.ToStringR(res))
				}
			}
		case "25":
			// Subtraction
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x25")
			} else {
				r2, err := InputRational(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x25")
				} else {
					res := rational.Subtraction(r1, r2)
					result.SetText(rational.ToStringR(res))
				}
			}
		case "26":
			// Multiplication
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x26")
			} else {
				r2, err := InputRational(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x26")
				} else {
					res := rational.Multiplication(r1, r2)
					result.SetText(rational.ToStringR(res))
				}
			}
		case "27":
			// Division
			r1, err := InputRational(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x27")
			} else {
				r2, err := InputRational(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x27")
				} else if entry2.Text == "" || entry2.Text == "0" {
					result.SetText("Ошибка нулевого делителя: x27")
				} else {
					res := rational.Division(r1, r2)
					result.SetText(rational.ToStringR(res))
				}
			}
		case "30":
			// AdditionP
			p1, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x30")
			} else {
				p2, err := InputPolynomial(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x30")
				} else {
					res := polynome.AdditionP(p1, p2)
					result.SetText(res.ToStringPol())
				}
			}
		case "31":
			// SubtractionP
			p1, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x31")
			} else {
				p2, err := InputPolynomial(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x31")
				} else {
					res := polynome.SubtractionP(p1, p2)
					result.SetText(res.ToStringPol())
				}
			}
		case "32":
			// MultiplicationRational
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x32")
			} else {
				r, err := InputRational(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x32")
				} else {
					res := polynome.MultiplicationRational(p, r)
					result.SetText(res.ToStringPol())
				}
			}
		case "33":
			// MultiplicationXpowerK
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x33")
			} else {
				k, _ := strconv.Atoi(entry2.Text)
				res := polynome.MultiplicationXpowerK(p, k)
				result.SetText(res.ToStringPol())
			}
		case "34":
			// OlderCoeffPoly
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x30")
			} else {
				result.SetText(rational.ToStringR(polynome.OlderCoeffPoly(p)))
			}
		case "35":
			// OlderPoly
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x30")
			} else {
				result.SetText(strconv.Itoa(int(polynome.OlderPoly(p))))
			}
		case "36":
			// Greatest... (what the hell?)
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x30")
			} else {
				n1, n2 := polynome.GreatestCommonDivisorAndLeastCommonMultipleOfPolynomial(p)
				result.SetText("НОД: " + natural.ToStringN(n1) + ", " + "НОК: " + natural.ToStringN(n2))
			}
		case "37":
			// MultiplicationPol
			p1, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x37")
			} else {
				p2, err := InputPolynomial(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x37")
				} else {
					res := polynome.MultiplicationPol(p1, p2)
					result.SetText(res.ToStringPol())
				}
			}
		case "38":
			// QuotientOfDivision
			p1, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x38")
			} else {
				p2, err := InputPolynomial(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x38")
				} else {
					if entry2.Text == "0" {
						result.SetText("Ошибка нулевого делителя: x38")
					} else {
						res := polynome.QuotientOfDivision(p1, p2)
						result.SetText(res.ToStringPol())
					}
				}
			}
		case "39":
			// RemainderFromDivision
			p1, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x39")
			} else {
				p2, err := InputPolynomial(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x39")
				} else {
					if entry2.Text == "0" {
						result.SetText("Ошибка нулевого делителя: x39")
					} else {
						res := polynome.RemainderFromDivision(p1, p2)
						result.SetText(res.ToStringPol())
					}
				}
			}
		case "3A":
			// GreatestCommonDivisor
			p1, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x3A")
			} else {
				p2, err := InputPolynomial(entry2.Text)
				if err {
					result.SetText("Ошибка нулевого знаменателя: x3A")
				} else {
					res := polynome.GreatestCommonDivisor(p1, p2)
					result.SetText(res.ToStringPol())
				}
			}
		case "3B":
			// Derivative
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x3B")
			} else {
				res := polynome.Derivative(p)
				result.SetText(res.ToStringPol())
			}
		case "3C":
			// SimplifyRoots
			p, err := InputPolynomial(entry1.Text)
			if err {
				result.SetText("Ошибка нулевого знаменателя: x3C")
			} else {
				res := polynome.SimplifyRoots(p)
				result.SetText(res.ToStringPol())
			}
		}
	})
	btnRes.Hide()
	// Result box
	resW := container.NewVBox(format, entry1, entry2, entry3, result, btnRes)

	btnNextFunc := widget.NewButton("К вводу", func() {
		switch Selected {
		case "0":
			switch funcBoxN.Selected {
			case funcsN[0]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsN[1]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsN[2]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsN[3]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsN[4]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsN[5]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsN[6]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsN[7]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
			case funcsN[8]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Show()
				btnRes.Show()
				SelectedFun = Selected + "8"
			case funcsN[9]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "9"
			case funcsN[10]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "A"
			case funcsN[11]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "B"
			case funcsN[12]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "C"
			case funcsN[13]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "D"
			}
		case "1":
			switch funcBoxWh.Selected {
			case funcsWh[0]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsWh[1]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsWh[2]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsWh[3]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsWh[4]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsWh[5]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsWh[6]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsWh[7]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
			case funcsWh[8]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "8"
			case funcsWh[9]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "9"
			}
		case "2":
			switch funcBoxR.Selected {
			case funcsR[0]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsR[1]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsR[2]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsR[3]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsR[4]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsR[5]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsR[6]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsR[7]:
				format.Hide()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
			}
		case "3":
			switch funcBoxPol.Selected {
			case funcsPol[0]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsPol[1]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsPol[2]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsPol[3]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsPol[4]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsPol[5]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsPol[6]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsPol[7]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
			case funcsPol[8]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "8"

			case funcsPol[9]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "9"
			case funcsPol[10]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "A"
			case funcsPol[11]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "B"
			case funcsPol[12]:
				format.Show()
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "C"
			}
		}

	})

	funcBox := container.NewVBox()
	cardFunc := widget.NewCard("Функции", "", funcBox)
	btnNextMod := widget.NewButton("К функциям", func() {
		switch radioMods.Selected {
		case "Натуральные":
			Selected = "0"
			funcBox = container.NewVBox(funcBoxN, btnNextFunc)
			cardFunc.SetContent(funcBox)
		case "Целые":
			Selected = "1"
			funcBox = container.NewVBox(funcBoxWh, btnNextFunc)
			cardFunc.SetContent(funcBox)
		case "Рациональные":
			Selected = "2"
			funcBox = container.NewVBox(funcBoxR, btnNextFunc)
			cardFunc.SetContent(funcBox)
		case "Полиномы":
			Selected = "3"
			funcBox = container.NewVBox(funcBoxPol, btnNextFunc)
			cardFunc.SetContent(funcBox)
		}
	})
	modBox := container.NewVBox(radioMods, btnNextMod)

	cardMod := widget.NewCard("Модули", "", modBox)

	cardRes := widget.NewCard("Результат", "", resW)
	w.SetContent(container.NewHBox(
		cardMod,
		cardRes,
		cardFunc,
	))
	w.ShowAndRun()
}

func InputNatural(s string) natural.Natural {
	if s == "" {
		s = "0"
	}
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

func InputWhole(s string) whole.Whole {
	var negative = false
	if s == "" {
		s = "0"
	}
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	nums := InputNatural(s).Digits
	var w whole.Whole
	w.MakeW(negative, nums)
	return w
}

func InputRational(s string) (rational.Rational, bool) {
	if s == "" {
		s = "0"
	}
	arr := strings.Split(s, "/")
	if len(arr) == 1 {
		arr = append(arr, "1")
	}
	if arr[1] == "0" {
		return rational.Zero(), true
	}
	nom := InputWhole(arr[0])
	denom := InputNatural(arr[1])
	var r rational.Rational
	r.MakeR(nom, denom)
	return r, false
}

func InputPolynomial(s string) (polynome.Polynomial, bool) {
	var arr1 []string
	var res polynome.Polynomial
	var NextNegative, ThisNegative bool
	var err bool
	if s == "" {
		return polynome.Polynomial{Older: 0, Coeff: []rational.Rational{rational.Zero()}}, false
	}
	res.Coeff = append([]rational.Rational{rational.Zero()}, res.Coeff...)

	// Split string on + -
	arr0 := strings.SplitAfter(s, "+")
	for i := 0; i < len(arr0); i++ {
		a := strings.SplitAfter(arr0[i], "-")
		for j := 0; j < len(a); j++ {
			arr1 = append(arr1, a[j])
		}
	}
	// Handle Case with one number
	if len(arr0) == 1 && strings.Count(arr0[0], "x") == 0 {
		res.Older = 0
		temp, err := InputRational(arr0[0])
		if err {
			return polynome.Polynomial{Older: 0, Coeff: []rational.Rational{rational.Zero()}}, true
		}
		res.Coeff = []rational.Rational{temp}
		return res, false
	}
	// Determine first sign
	if arr1[0][0] == '-' {
		NextNegative = true
		arr1[0] = arr1[0][1:]
	}
	if arr1[0] == "" {
		arr1 = arr1[1:]
	}
	for i := 0; i < len(arr1); i++ {
		a := strings.Split(arr1[i], "x^")
		// If no coeff -> coeff = 1/1
		if a[0] == "" {
			a[0] = "1/1"
		}

		// NextNegative was determined in previous loop
		ThisNegative = NextNegative
		NextNegative = false
		if len(a) == 1 {
			if strings.Count(a[0], "x") != 0 {
				// x^1 case
				// Deleting sign char
				if a[0][len(a[0])-1] == '-' {
					NextNegative = true
					a[0] = a[0][0 : len(a[0])-1]
				}
				if a[0][len(a[0])-1] == '+' {
					a[0] = a[0][0 : len(a[0])-1]
				}
				// Deleting 'x'
				a[0] = a[0][0 : len(a[0])-1]
				// Handle 'x' case
				if a[0] == "" {
					a[0] = "1/1"
				}
				if res.Older < 1 {
					res.Older = 1
					for uint32(len(res.Coeff)) <= 1 {
						res.Coeff = append([]rational.Rational{rational.Zero()}, res.Coeff...)
					}
				}
				res.Coeff[res.Older-1], err = InputRational(a[0])
				if err {
					return polynome.Polynomial{Older: 0, Coeff: []rational.Rational{rational.Zero()}}, true
				}
				res.Coeff[res.Older-1].Nominator.Negative = ThisNegative
			} else {
				res.Coeff[res.Older], err = InputRational(a[0])
				if err {
					return polynome.Polynomial{Older: 0, Coeff: []rational.Rational{rational.Zero()}}, true
				}
				res.Coeff[res.Older].Nominator.Negative = ThisNegative
			}
		} else {
			// Determine next sign
			if a[1][len(a[1])-1] == '-' {
				NextNegative = true
				a[1] = a[1][0 : len(a[1])-1]
			} else if a[1][len(a[1])-1] == '+' {
				a[1] = a[1][0 : len(a[1])-1]
			}
			// Power of this
			powI, _ := strconv.Atoi(a[1])
			pow := uint32(powI)
			// Power is more that we have now -> we need to expand coeff array
			if res.Older < pow {
				res.Older = pow
				for uint32(len(res.Coeff)) <= pow {
					res.Coeff = append([]rational.Rational{rational.Zero()}, res.Coeff...)
				}
			}
			// Handling Coeff
			res.Coeff[res.Older-pow], err = InputRational(a[0])
			if err {
				return polynome.Polynomial{Older: 0, Coeff: []rational.Rational{rational.Zero()}}, true
			}
			res.Coeff[res.Older-pow].Nominator.Negative = ThisNegative
		}
	}
	return res, false
}
