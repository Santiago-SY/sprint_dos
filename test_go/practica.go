package main

type LinkedListNode struct {
	element int
	next    *LinkedListNode
}

func reverse(list *LinkedListNode) *LinkedListNode {
	var inicio *LinkedListNode = nil
	for list != nil {
		nuevo := &LinkedListNode{}
		nuevo.element = list.element
		nuevo.next = inicio
		inicio = nuevo
		list = list.next
	}
	return inicio
}

func symmetricDifference(firstList *LinkedListNode, secondList *LinkedListNode) *LinkedListNode {
	conteo := make(map[int]int)

	for firstList != nil {
		conteo[firstList.element]++
		firstList = firstList.next
	}
	for secondList != nil {
		conteo[secondList.element]++
		secondList = secondList.next
	}

	var resultado *LinkedListNode = nil

	for numero, cantidad := range conteo {

		if cantidad == 1 {

			nuevo := &LinkedListNode{}
			nuevo.element = numero // Guardamos el numero (la clave), no la cantidad

			nuevo.next = resultado
			resultado = nuevo
		}
	}

	return resultado
}

func binarySearch(sorted []int, element int) int {
	var inicio int = 0
	var final int = len(sorted) - 1
	for inicio <= final {
		var medio int = (inicio + final) / 2
		if sorted[medio] == element {
			return medio
		}

		if sorted[medio] < element {
			inicio = medio + 1
			//codigo para que recorra desde el medio hacia el final
		}
		if sorted[medio] > element {
			final = medio - 1
			//codigo para que recorra desde el inicio hasta el medio que ahora seria el final
		}
	}
	return -1
}

/**
 * Count the number of duplicate array elements.
 * Duplicate is defined as two or more identical elements.
 * @param numbers is an array of int
 * @return the amount of duplicate numbers
 * For example, in the array [1, 2, 2, 3, 3, 3], the two twos are one duplicate and so are the three threes.
 * It returns 2.
 */
func countDuplicates(numbers []int) int {
	var recorre int = 0
	conteo := make(map[int]int)
	var res int = 0
	for recorre < len(numbers) {
		conteo[numbers[recorre]]++
		recorre++
	}
	for _, cantidad := range conteo {

		if cantidad > 1 {

			res++
		}
	}
	println("la cantidad de repetidos es:")
	return res
}
