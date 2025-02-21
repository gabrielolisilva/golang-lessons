package main

import (
	"fmt"
	"golangLessons/testpackage"
)

/* Golang has only one main function in folder scope */
func main() {
	testFunction();
	testpackage.PackageFunction();

	var n1 int = 10;
	var n2 int = 20;
	testpackage.SumNumbers(n1, n2);

	sum, multiply := testpackage.SumAndMultiplyValues(n1, n2);
	fmt.Println(sum, multiply);

	for x := 0; x < 10; x++ {
		fmt.Println(x);
	}
}

func testFunction() {
	fmt.Println("Step 1");
	fmt.Println("Step 2");
}