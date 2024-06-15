package main

import "fmt"

func isValidOperator(operator, operators string) bool {
	for _, value := range operators {
		if operator == string(value) {
			return true
		}
	}
	return false
}

func main() {

	var value1 float64
	var value2 float64
	var operator string

	var operators string = "+-*/"

	fmt.Println("Enter two values for calculations")

	fmt.Scanf("%f %f", &value1, &value2)

	fmt.Println("Enter the operator")
	fmt.Scanf("%s", &operator)

	valid := isValidOperator(operator, operators)

	if !valid {
		fmt.Println("The provided operator is not applicable")
	}

	switch operator {
	case "+":
		fmt.Printf("The Sum of %0.2f + %0.2f Numbers is %0.2ff\n", value1, value2, value1+value2)
	case "-":
		fmt.Printf("The Difference of %0.2f - %0.2f Numbers is %0.2f\n", value1, value2, value1-value2)
	case "*":
		fmt.Printf("The Product of  Numbers  %0.2f - %0.2f  is %0.2f\n", value1, value2, value1*value2)
	case "/":
		if value2 == 0 {
			fmt.Println("Divide by zero is not possible")
		} else {
			fmt.Printf("The Product of  %0.2f / %0.2f  Numbers is %0.2f\n", value1, value2, value1/value2)
		}
	}

}
