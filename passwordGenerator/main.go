package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length int
	fmt.Println("Enter the size of password")
	fmt.Scanf("%d", &length)

	includeNumbers := flag.Bool("number", false, "Do you want to include numbers yes or no")
	includeSymbols := flag.Bool("symbol", false, "Do you want to include numbers yes or no")
	flag.Parse()

	password := make([]byte, length)

	total := 2

	if *includeNumbers {
		total += 1
	}

	if *includeSymbols {
		total += 1
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		rand1 := rand.Intn(total)

		switch rand1 {
		case 0:
			password[i] = byte(rand.Intn(26) + 'a')
		case 1:
			password[i] = byte(rand.Intn(26) + 'A')
		case 2:
			password[i] = byte(rand.Intn(10) + '0')
		case 3:
			symbols := "!@#$%^&*()_+"
			password[i] = byte(symbols[rand.Intn(len(symbols))])
		}

	}
	fmt.Println("The Password is", string(password))
}
