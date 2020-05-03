package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter a number between 1 and 5: ")
	var file *os.File
	file = os.Stdin
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	value, err := strconv.ParseInt(scanner.Text(), 10, 64)

	//An if conditional statement
	if err != nil {
		fmt.Println("Error: value is not a number")
		return
	}

	// An if statement with statement and conditional
	// v has a scope within the if statement only
	if v := value; v > 5 || v < 1 {
		fmt.Println("The value is not within the range 1 to 5")
		return
	}

	// Switch statement with single value conditons

	fmt.Print("This finger number represents: ")
	switch value {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}


	// Switch with initial statement and multiple case values
	switch letter := "i"; letter{
	case "a", "e", "i", "o", "u":
		fmt.Println("These are vowels")
	default:
		fmt.Println("Letter is not a vowel")
	}

	// Switch without expressions and the fallthrough statement
	number := 20
	switch{
	case number < 10:
		fmt.Println("Number is less than 10")
		fallthrough
	case number < 20:
		fmt.Println("Number is less than 20")
		fallthrough
	case number < 50:
		fmt.Println("Number is less than 50")
		fallthrough
	case number < 100:
		fmt.Println("Number is less than 100")
	}
}
