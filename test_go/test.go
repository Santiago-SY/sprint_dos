package main

import (
	"fmt"
	"math"
	"time"
)

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	/*
		The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
		This syntax is only available inside functions.
	*/
	//var f string = "apple"
	f := "apple"
	fmt.Println(f)
}

// parte de constantes
const s string = "constant"

/*
Las constantes son valores de precisión arbitraria que infieren su tipo según el contexto
(si la operación pide int, se vuelve int; si pide float, se vuelve float)
*/
func constants() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func for_test() {
	i := 1
	//esto seria el while en c++
	// En C++ sería: while (i <= 3) { ... }
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//for normal
	// En C++ sería: for(int j=0; j<3; j++) { ... }
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// TIPO 3: El "Range" (Específico de Go)
	// Es una forma rápida de decir "Repetir N veces"
	// (Empieza en 0 y termina antes de llegar al 3)
	for i := range 3 {
		fmt.Println("range", i)
	}
	// TIPO 4: El "While True" (Bucle Infinito)
	// En C++ sería: while(true) { ... }
	// Se usa mucho en servidores o procesos que no deben parar nunca
	for {
		fmt.Println("loop")
		break
	}
	// TIPO 5: El "Foreach" con lógica extra
	// Itera sobre el rango del 0 al 5
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func if_else_test() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("either 8 or 7 are even")
	}
	/*
		inicializa num := 0 (crea la variable ahi mismo)
		evalua num < 0
	*/
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switch_test() {
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("im am bool")
		case int:
			fmt.Println("im an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func main() {
	//fmt.Println("hello world")
	//values()
	//variables()
	//constants()
	//for_test()
	//if_else_test()
	switch_test()
}
