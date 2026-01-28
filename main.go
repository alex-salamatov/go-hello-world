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

	if !scanner.Scan() {
		fmt.Println("Failed to read age")
		return
	}
	age, err := strconv.Atoi(scanner.Text())
	birthYearString := "unknown"
	if err != nil {
		fmt.Println("Entered age is not valid")
	}

	fmt.Printf("Hello %s, your year of birth is ", name)
	if isAgeOk(age) {
		birthYear = birthYear - age
		birthYearString = fmt.Sprintf("%d or %d\n", birthYear, birthYear+1) //you never know actual year of birth without knowing exact birthday
	}
	fmt.Println(birthYearString)
}
