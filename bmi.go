package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Output information
	fmt.Println("BMI Calculator")
	fmt.Println("--------------------")

	// prompt for user input (weight + height)
	fmt.Print("Please enter your wight (kg): ")
	weightInput, _ := reader.ReadString('\n') // read until the user hits enter

	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n') // read until the user hits enter

	// save user input in variables
	if runtime.GOOS == "windows" {
		weightInput = strings.Replace(weightInput, "\r\n", "", -1)
		heightInput = strings.Replace(heightInput, "\r\n", "", -1)
	} else {
		weightInput = strings.Replace(weightInput, "\n", "", -1)
		heightInput = strings.Replace(heightInput, "\n", "", -1)
	}
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	// calculate BMI (weight / height * height)
	bmi := weight / (height * height)
	// output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
