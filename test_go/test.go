package main

/*
Arrays
Slices
Maps
Range
Functions
GoROutines
Channels
Channel Syncronization
Mutexes
Pointers
Struct
*/

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"sync"
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

func array_test() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func slices_test() {
	//Corchetes vacios = esto puede crecer o acihcarse,es dinamico
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	//el constructor, 1)reserva memoria para 3 strings, 2)los pone en 0 o vacio
	//3)te retorna el slice listo para usar
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	//Append agrega elementos al final de la lista
	//Si no hay espacio, go automaticamente duplica la memoria por debajo
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // crea un slice vacio del mismo tamaño
	//Copy es clonar
	//como los slices son ventanas a un array en memoria,si no se usa copy y solo se hace c:= s
	//modificar c tambien cambiaria  s, con Copy son independientes
	copy(c, s) //copias s dentro de c
	fmt.Println("cpy:", c)
	//empieza en indice 2 (c)
	//termina antes de 5(toma hasta el 4 que es e)
	//la regla matematica es exclusive/inclusive
	//incluye el primer numero
	//excluye el segundo numero
	//s[:5] = Desde el principio hasta el 4.
	//s[2:] = Desde el 2 hasta el final.
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		Para crear la lista: lista := []int{1, 2, 3} o make([]int, 0).

		Para agregar cosas: lista = append(lista, valor).

		Para recorrerla: for i, v := range lista.

		Para saber el tamaño: len(lista).
	*/
}

func maps_test() {
	//       Clave      Valor
	//String: las llaves seran textos
	//Int: los valores seran numericos
	// al igual que las slices,se inicializan con makes
	//Si solo declaras var m map[string]int, el mapa es nil y si intentas escribir en él,
	// el programa explota (panic).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	//Si buscas una clave que NO EXISTE, Go no da error. Te devuelve el "valor cero" del tipo de dato.
	v3 := m["k3"]          //k3 no existe en el mapa
	fmt.Println("v3:", v3) // imprime 0 (porque es un int) si el mapa fuera de string devuelve "" o bool false

	fmt.Println("len:", len(m))
	//delete (m, "clave") Borra una entrada especifica, si la clave no existe, no hace nada
	//no da error
	//clear (m) Borra todo el mapa, dejandolo vacio
	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)
	//Cuando accedes a un mapa se puede pedir dos valores de retorno
	//el valor (en este caso se ignora con _)
	// el booleano (prs o ok): te dice true si la clave existe, y sino false
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func range_test() {

	nums := []int{2, 3, 4}
	sum := 0
	//range devuelve dos cosas indice, valor
	//usa _ para que no traiga el indice porque no se va a usar
	//si trajera i en lugar de _ y no lo uso,el compilador tira error
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//i es la posicion (0, 1 y 2) y num es el valor
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	//range con maps (clave y valor)
	//cuando recorro un mapa range devuelve
	//k: la clave ej "a", "b"
	//v el valor ej "apple", "banana"
	for k, v := range kvs {
		//al imprimir usa Printf (Print con Formato) y %s. Es como en C (printf).
		//El %s en la función fmt.Printf de Go es un especificador de formato
		//que se utiliza para imprimir un valor como una cadena de texto (string).
		fmt.Printf("%s -> %s\n", k, v)
	}
	//Si solo pones una variable antes del := al recorrer un mapa, Go asume que solo quieres las Claves.
	//Los valores los ignora.
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range con strings
	//imprime 103 y 111 porque go recorre el string letra por letra
	//y c es el valor ASCII de la letra
	//g = 103 o = 111
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

//funcion para goroutines

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func goroutines_test() {
	//llamada normal sincrona
	//el programa principal llega a esta linea y se detiene
	//entra a la funcion, imprime 0,1,2
	//recien cuando termina,el programa sigue linea abajo
	//por eso el output es ordenado
	f("direct")

	//llamada normal asincrona
	//al poner la palabra go antes de la funcion, lo lanza en un hilo a parte y no espera a que termine
	//el programa no se detiene, Salta a la otra linea
	//mientras tanto en paralelo, empieza la funcion f
	go f("goroutine")

	//la funcion minima
	//Es lo mismo que la anterior, pero definiendo la función ahí mismo (como una lambda).
	//También arranca en paralelo.
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	/*La funcion main es el hilo principal, si main termina, el programa muere
	y mata a todos los goroutines hijos hayan terminado o no
	tira un sleep de un segundo para darle tiempo a los goroutines de terminar*/
	time.Sleep(time.Second)
	fmt.Println("done")

}

func Channel_test() {
	//chan string es una tuberia tipada
	//por este tubo solo pueden viajar strings
	//si intentas meter un numero,explota (error de compilacion)
	//se crea con make
	messages := make(chan string)

	go func() {
		// el <- indica a donde viajan los datos
		//enviar(meter en el tubo)
		//mete la palabra ping dentro de la variable messages
		//la flecha apunta hacia el canal
		messages <- "ping"
	}()
	//recibir (sacar del tubo)
	//lectura Saca un dato desde messages y lo guarda en msg
	//la flecha sale del canal
	/*el main llega a esta linea y como el canal esta vacio
	se congela en esa linea
	mientras tanto el go routine que corre en paralelo mete "ping" en el tubo
	cuando llega el dato,el main se descongela,agarra el dato,lo imprime y termina
	El canal actúa como un semáforo automático. El que recibe espera al que envía.*/
	msg := <-messages
	fmt.Println(msg)
}

//func para el channel_syncro

func worker(done chan bool) {
	// la funcion recibe un canal para avisar
	fmt.Print("working...")
	time.Sleep(time.Second) //simula trabajo
	fmt.Println("done")
	//avisa mandando un true al tube para decir que termino
	//No importa el valor: Podrías mandar true, false, 1 o "hola".
	// Lo único que importa es el hecho de recibir algo.
	// El acto de recibir es lo que desbloquea al main.
	done <- true
}

func channel_syncro_test() {
	//creamos el canal de notificacion
	done := make(chan bool, 1)
	//lanza al trabajador y le da el canal
	go worker(done)
	//bloqueo
	// Esta línea se queda congelada esperando a que salga ALGO del tubo.
	<-done

}

//func para el mutexes

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock() // el candado- cierra la puerta
	/*defer significa: "Ejecuta esta línea JUSTO antes de que termine la función".
	Es una buena práctica poner el Unlock con defer inmediatamente después del Lock.
	Así te aseguras de nunca olvidarte de abrir el candado
	(lo que causaría un Deadlock y congelaría el programa).*/
	defer c.mu.Unlock() //acuerdate de escribir la salida
	c.counters[name]++  //el recurso compartido - hace el cambio
}

func mutexes_test() {

	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("b", 10000)
	}()

	wg.Wait()
	fmt.Println(c.counters)
}

//func para pointers_test

func zeroval(ival int) {
	//pasado por copia
	ival = 0
}

func zeroptr(iptr *int) {
	//pasado por referencia
	//*int Significa "No quiero un número, quiero la dirección de memoria donde vive un número
	//Convierte un puntero (*int) en el valor normal (int) para leerlo o escribirlo.
	*iptr = 0
}

func pointers_test() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	//Convierte una variable normal (int) en un puntero (*int).
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}
func main() {
	//fmt.Println("hello world")
	//values()
	//variables()
	//constants()
	//for_test()
	//if_else_test()
	//switch_test()
	//array_test()
	//slices_test()
	//maps_test()
	//range_test()
	//goroutines_test()
	//Channel_test()
	//channel_syncro_test()
	mutexes_test()
	pointers_test()

}
