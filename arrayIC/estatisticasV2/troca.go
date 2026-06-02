package main

func troca(index1 int, index2 int, array []int) {
	var temp int

	temp = array[index1]
	array[index1] = array[index2]
	array[index2] = temp
}
