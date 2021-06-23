package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// can't be a const compile don't run functions
// even though once function runs it won't change again that function is only
// ran at runtime
var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const seperator = "--------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightPrompt = "Please enter your height (m): "

func main() {
	// Output information
	fmt.Println(mainTitle)
	fmt.Println(seperator)

	// prompt for user input (weight + height)
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n') // read until the user hits enter

	fmt.Print(heightPrompt)
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
