package main

// GUI Михаил Голубев & Максим Тростин
import (
	"DMTI_baza/natural"
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

	entry1.Hide()
	entry2.Hide()
	entry3.Hide()
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
			res := natural.DivideOneIteration(n1, n2)
			result.SetText(natural.ToStringN(res))
		case "0A":
			// IntegerFromDivision
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.IntegerFromDivision(n1, n2)
			result.SetText(natural.ToStringN(res))
		case "0B":
			// RemainderFromDivision
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.RemainderFromDivision(n1, n2)
			result.SetText(natural.ToStringN(res))
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
			n1 := InputNatural(entry1.Text)
			res := whole.FromNaturalsToWhole(n1)
			result.SetText(whole.ToStringW(res))
		case "14":
			w1 := InputWhole(entry1.Text)
			res := whole.FromWholeToNaturals(w1)
			result.SetText(natural.ToStringN(res))
		case "15":
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.Addition(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "16":
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.Subtraction(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "17":
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.Multiplication(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "18":
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.WholeFromDivision(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "19":
			w1 := InputWhole(entry1.Text)
			w2 := InputWhole(entry2.Text)
			res := whole.RemainderFromDivision(w1, w2)
			result.SetText(whole.ToStringW(res))
		case "20":
			r1 := InputRational(entry1.Text)
			res := rational.SimplifyingFractions(r1)
			result.SetText(rational.ToStringR(res))
		case "21":
			r1 := InputRational(entry1.Text)
			res := rational.CheckingForWhole(r1)
			if res {
				result.SetText("Да")
			} else {
				result.SetText("Нет")
			}
		case "22":
			w1 := InputWhole(entry1.Text)
			res := rational.WholeToFractional(w1)
			result.SetText(rational.ToStringR(res))
		case "23":
			r1 := InputRational(entry1.Text)
			res := rational.FractionalToWhole(r1)
			result.SetText(whole.ToStringW(res))
		case "24":
			r1 := InputRational(entry1.Text)
			r2 := InputRational(entry2.Text)
			res := rational.Addition(r1, r2)
			result.SetText(rational.ToStringR(res))
		case "25":
			r1 := InputRational(entry1.Text)
			r2 := InputRational(entry2.Text)
			res := rational.Subtraction(r1, r2)
			result.SetText(rational.ToStringR(res))
		case "26":
			r1 := InputRational(entry1.Text)
			r2 := InputRational(entry2.Text)
			res := rational.Multiplication(r1, r2)
			result.SetText(rational.ToStringR(res))
		case "27":
			r1 := InputRational(entry1.Text)
			r2 := InputRational(entry2.Text)
			res := rational.Division(r1, r2)
			result.SetText(rational.ToStringR(res))
		}
	})
	btnRes.Hide()
	//result
	resW := container.NewVBox(entry1, entry2, entry3, result, btnRes)

	btnNextFunc := widget.NewButton("К вводу", func() {
		switch Selected {
		case "0":
			switch funcBoxN.Selected {
			case funcsN[0]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsN[1]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsN[2]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsN[3]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsN[4]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsN[5]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsN[6]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsN[7]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
			case funcsN[8]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Show()
				btnRes.Show()
				SelectedFun = Selected + "8"
			case funcsN[9]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "9"
			case funcsN[10]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "A"
			case funcsN[11]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "B"
			case funcsN[12]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "C"
			case funcsN[13]:
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
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsWh[1]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsWh[2]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsWh[3]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsWh[4]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsWh[5]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsWh[6]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsWh[7]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
			case funcsWh[8]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "8"
			case funcsWh[9]:
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
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "0"
			case funcsR[1]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "1"
			case funcsR[2]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "2"
			case funcsR[3]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Hide()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "3"
			case funcsR[4]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "4"
			case funcsR[5]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "5"
			case funcsR[6]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "6"
			case funcsR[7]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				entry1.Show()
				entry2.Show()
				entry3.Hide()
				btnRes.Show()
				SelectedFun = Selected + "7"
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
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	nums := InputNatural(s).Digits
	var w whole.Whole
	w.MakeW(negative, nums)
	return w
}

func InputRational(s string) rational.Rational {
	arr := strings.Split(s, "/")
	nom := InputWhole(arr[0])
	denom := InputNatural(arr[1])
	var r rational.Rational
	r.MakeR(nom, denom)
	return r
}
