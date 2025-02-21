package main

import (
	"fmt"
	"golangLessons/testpackage"
)

/* Golang has only one main function in folder scope */
func main() {
	testFunction(2);
	testpackage.PackageFunction();

	var n1 int = 10;
	var n2 int = 20;
	testpackage.SumNumbers(n1, n2);

	sum, multiply := testpackage.SumAndMultiplyValues(n1, n2);
	fmt.Println(sum, multiply);

	for x := 0; x < 10; x++ {
		fmt.Println(x);
	}

	var agesIlimited [3]int = [3]int{1, 2, 3};
	agesLimited := []string{"a", "b", "c", "d"};
	agesLimited = append(agesLimited, "e");

	fmt.Println(agesIlimited, len(agesIlimited), agesLimited, len(agesLimited));

	var slicedAges []string = agesLimited[1:3];
	var slicedAgesToEnd []string = agesLimited[1:];
	var slicedAgesBegginingToElement []string = agesLimited[:3];
	fmt.Println(slicedAges, slicedAgesToEnd, slicedAgesBegginingToElement);
}

func testFunction(value int) {
	if value == 1 {
		fmt.Println("Step 1");
	} else {
		fmt.Println("Step 2");
	}
}