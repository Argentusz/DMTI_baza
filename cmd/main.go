package main

//Исполняемоый пакет для работы с консольными флагами
//В целом, на него смотреть пока не нужно, это я сам добью, потом посмотрите и оцените(!)
import (
	"DMTI_baza/polynome"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//Модуль от Голубева Михаила
func main() {
	var moduleNumber int
	var functionNumber int
	flag.IntVar(&moduleNumber, "m", 0, "name of module")     //флаг выбора модуля от 0 до 3, по умолчанию модуль 0
	flag.IntVar(&functionNumber, "f", 0, "name of function") //флаг выбора номер функции, по умолчанию - 0
	flag.Parse()

	if moduleNumber == 3 {
		polOperations(functionNumber)
	} else {
		fmt.Println("не полиномы")
	}
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func polOperations(funcNo int) {
	var p1, p2 polynome.Polynome
	var result polynome.Polynome
	switch funcNo {
	case 0:
		fmt.Println("Это функция сложения полиномов, добро пожаловать")
		var arrP1 []float64
		var arrP2 []float64
		var count1, count2 int
		var val1, val2 float64
		fmt.Println("Введите старшую степень первого полинома:")
		fmt.Scan(&count1)
		for i := 0; i <= count1; i++ {
			fmt.Printf("Введите коэффициент при степени %d :\n", count1-i)
			fmt.Scan(&val1)
			arrP1 = append(arrP1, val1)
		}
		fmt.Println("Введите старшую степень второго полинома:")
		fmt.Scan(&count2)
		for i := 0; i <= count2; i++ {
			fmt.Printf("Введите коэффициент при степени %d :\n", count2-i)
			fmt.Scan(&val2)
			arrP2 = append(arrP2, val2)
		}
		p1.MakePol(arrP1)
		p2.MakePol(arrP2)
		CallClear()
		fmt.Println("Первый полином:", p1.ToString())
		fmt.Println("Второй полином:", p2.ToString())
		result = polynome.AdditionPol(p1, p2)
		fmt.Println("Результат сложения:", result.ToString())
	case 1:
		fmt.Println("Это функция вычиитания полиномов, добро пожаловать")
		var arrP1 []float64
		var arrP2 []float64
		var count1, count2 int
		var val1, val2 float64
		fmt.Println("Введите старшую степень первого полинома:")
		fmt.Scan(&count1)
		for i := 0; i <= count1; i++ {
			fmt.Printf("Введите коэффициент при степени %d :\n", count1-i)
			fmt.Scan(&val1)
			arrP1 = append(arrP1, val1)
		}
		fmt.Println("Введите старшую степень второго полинома:")
		fmt.Scan(&count2)
		for i := 0; i <= count2; i++ {
			fmt.Printf("Введите коэффициент при степени %d :\n", count2-i)
			fmt.Scan(&val2)
			arrP2 = append(arrP2, val2)
		}
		p1.MakePol(arrP1)
		p2.MakePol(arrP2)
		CallClear()
		fmt.Println("Первый полином:", p1.ToString())
		fmt.Println("Второй полином:", p2.ToString())
		result = polynome.SubtractionPol(p1, p2)
		fmt.Println("Результат вычитания:", result.ToString())
	case 2:
		fmt.Println("Это функция умножения полинома на рациональное число, добро пожаловать")
		var arrP1 []float64
		var count1 int
		var val1, val2 float64
		fmt.Println("Введите старшую степень полинома:")
		fmt.Scan(&count1)
		for i := 0; i <= count1; i++ {
			fmt.Printf("Введите коэффициент при степени %d :\n", count1-i)
			fmt.Scan(&val1)
			arrP1 = append(arrP1, val1)
		}
		fmt.Println("Введите число, на которое нужно умножить полином:")
		fmt.Scan(&val2)
		p1.MakePol(arrP1)
		result = polynome.NumberMultiplyPol(p1, val2)
		fmt.Println(result.ToString())
	case 3:
		fmt.Println("Это функция умножения полинома на x в степени k, добро пожаловать")
		var arrP1 []float64
		var count1, val2 int
		var val1 float64
		fmt.Println("Введите старшую степень полинома:")
		fmt.Scan(&count1)
		for i := 0; i <= count1; i++ {
			fmt.Printf("Введите коэффициент при степени %d :\n", count1-i)
			fmt.Scan(&val1)
			arrP1 = append(arrP1, val1)
		}
		fmt.Println("Введите k:")
		fmt.Scan(&val2)
		p1.MakePol(arrP1)
		result = polynome.VarMultiplyPol(p1, val2)
		fmt.Println(result.ToString())

	}
}
