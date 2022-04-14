package main

import (
	"DMTI_baza/natural"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func main() {

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

	a := app.New()
	w := a.NewWindow("bazalt")
	w.Resize(fyne.NewSize(1000, 500))

	///modules
	radioMods := widget.NewRadioGroup([]string{"Натуральные", "Целые", "Рациональные", "Полиномы"}, func(s string) {

	})
	radioMods.SetSelected("Натуральные")

	//natural
	funcBoxN := widget.NewRadioGroup(funcsN, func(s string) {})
	funcBoxN.SetSelected(funcsN[0])

	//whole
	funcBoxWh := widget.NewRadioGroup(funcsWh, func(s string) {})
	funcBoxWh.SetSelected(funcsWh[0])

	//rational
	funcBoxR := widget.NewRadioGroup(funcsR, func(s string) {})
	//funcBoxWh.SetSelected(funcsWh[0])

	//polynomial
	funcBoxPol := widget.NewRadioGroup(funcsPol, func(s string) {})

	//entries
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	Selected := ""
	SelectedFun := "0"
	result := widget.NewLabel("")

	btnRes := widget.NewButton("результат", func() {
		switch SelectedFun {
		case "00":
			n1 := InputNatural(entry1.Text)
			n2 := InputNatural(entry2.Text)
			res := natural.Compare(n1, n2)
			resW.Add(result)
			result.SetText(strconv.Itoa(res))
			resW = container.NewVBox(entry1)
		}
	})

	//result
	resW := container.NewVBox(entry1, entry2, btnRes)

	btnNextFunc := widget.NewButton("к вводу", func() {
		switch Selected {
		case "0":
			switch funcBoxN.Selected {
			case funcsN[0]:
				resW.Add(entry2)
				resW.Add(btnRes)
				SelectedFun = Selected + "0"
			case funcsN[1]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				resW.Remove(entry2)
				resW.Add(btnRes)
				SelectedFun = Selected + "1"
			case funcsN[2]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				resW.Remove(entry2)
				resW.Add(btnRes)
				SelectedFun = Selected + "2"
			case funcsN[3]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				resW.Add(entry2)
				resW.Add(btnRes)
				SelectedFun = Selected + "3"
			case funcsN[4]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				resW.Add(entry2)
				resW.Add(btnRes)
				SelectedFun = Selected + "4"
			case funcsN[5]:
				result.SetText("")
				entry1.SetText("")
				entry2.SetText("")
				resW.Add(entry2)
				resW.Add(btnRes)
				SelectedFun = Selected + "5"
			}
		}

	})

	funcBox := container.NewVBox()
	cardFunc := widget.NewCard("функции", "", funcBox)
	btnNextMod := widget.NewButton("к функциям", func() {
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

	cardMod := widget.NewCard("модули", "", modBox)

	cardRes := widget.NewCard("результат", "", resW)
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
