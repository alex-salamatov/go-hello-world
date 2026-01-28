package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const maxAge = 100 //assuming older people cannot work with computers :-)

func isAgeOk(age int) bool {
	if (age < 0) || (age > maxAge) {
		return false
	}

	return true
}

func main() {
	fmt.Print("Enter your name: ")

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Failed to read name")
		return
	}
	name := scanner.Text()

	fmt.Print("Enter your age: ")
	var age int
	var birthYear int = time.Now().Year()

	for {
		if !scanner.Scan() {
			fmt.Println("Failed to read age")
			return
		}

		var err error
		age, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Entered age is not a number, please try again to enter integer value")
			continue
		}

		if !isAgeOk(age) {
			fmt.Println("Entered age is not in valid range, please try again and enter en integer value in a range [0 - 100]")
			continue
		}

		break
	}

	birthYear = birthYear - age
	fmt.Printf("Hello %s, your year of birth is %d or %d\n", name, birthYear, birthYear+1) //you never know actual year of birth without knowing exact birthday
}
