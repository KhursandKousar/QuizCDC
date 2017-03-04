package main

import "fmt"

//RemoveFirstAndLast will remove the first and last element of the string array
func RemoveFirstAndLast(a []string) []string {
	newone := a[1 : len(a)-1]
	return newone //newone is a slice
}

func display(a []string, message string) {
	fmt.Print(message)
	for index := 0; index < len(a); index++ {
		fmt.Printf("%v ==>> %v\n", index+1, a[index])
	}

	fmt.Println()
}
func main() {
	strArrr := []string{"khursand", "kousar", "fastian", "concurrent", "sir"}

	display(strArrr, "\t\tThe array before removing \n")

	strArrr = RemoveFirstAndLast(strArrr)

	display(strArrr, "\tThe array after removing the first and second value \n")
}
