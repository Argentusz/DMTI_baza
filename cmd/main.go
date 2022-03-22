package main

import (
	"example.com/m/polynome"
	"flag"
	"fmt"
)

func main() {
	var counts int
	var funcs int
	var p1, p2 polynome.Polynome
	var pTemp polynome.Polynome
	flag.IntVar(&counts, "m", 0, "name of module")
	flag.IntVar(&funcs, "f", 0, "name of function")
	flag.Parse()

	if counts == 3 {
		if funcs == 0 {
			var arrP1 []float64
			var arrP2 []float64
			var count1, count2 int
			var val1, val2 float64
			fmt.Scan(&count1)
			for i := 0; i < count1; i++ {
				fmt.Scan(&val1)
				arrP1 = append(arrP1, val1)
			}

			fmt.Scan(&count2)
			for i := 0; i < count2; i++ {
				fmt.Scan(&val2)
				arrP1 = append(arrP1, val2)
			}
			p1.MakePol(arrP1)
			p2.MakePol(arrP2)
			pTemp = polynome.AdditionPol(p1, p2)
			fmt.Println(pTemp.ToString())
		}
	} else {
		fmt.Println("не полиномы")
	}
}
