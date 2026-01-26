package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const maxAge = 100 //assuming older people cannot work with computers :-)

func isAgeOk(age int, err error) bool {
	if err != nil {
		return false
	}

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

	if !scanner.Scan() {
		fmt.Println("Failed to read age")
		return
	}
	age, ageError := strconv.Atoi(scanner.Text())

	fmt.Printf("Hello %s, your year of birth is ", name)
	if isAgeOk(age, ageError) {
		birthYear = birthYear - age
		fmt.Printf("%d or %d\n", birthYear, birthYear+1) //you never know actual year of birth without knowing exact birthday
	} else {
		fmt.Println("unknown.")
	}

}
