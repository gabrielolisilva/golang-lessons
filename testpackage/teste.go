package testpackage

import "fmt"

/**
* If wants to export the function it has to be
* in package scope and capital letter
 */
func PackageFunction() {
	fmt.Println("Test function from testpackage");
}

func SumNumbers(n1, n2 int) {
	fmt.Println(n1 + n2);
}

func SumAndMultiplyValues(n1, n2 int) (sum, multiply int) {
	return n1 + n2, n1 * n2
}