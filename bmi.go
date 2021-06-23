package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/bangarangler/go-bmi/info"
)

func main() {
	// Output information
	fmt.Print(info.MesurmentType)
	mesurmentType, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		mesurmentType = strings.Replace(mesurmentType, "\r\n", "", -1)
		mesurmentType = strings.Replace(mesurmentType, "\r\n", "", -1)
	} else {
		mesurmentType = strings.Replace(mesurmentType, "\n", "", -1)
		mesurmentType = strings.Replace(mesurmentType, "\n", "", -1)
	}
	fmt.Println(info.MainTitle)
	fmt.Println(info.Seperator)

	// prompt for user input (weight + height)
	switch mesurmentType {
	case "metric":
		fmt.Print(info.WeightPromptM)
		weightInput, _ := reader.ReadString('\n') // read until the user hits enter

		fmt.Print(info.HeightPromptM)
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
	case "english":
		fmt.Print(info.WeightPromptE)
		weightInput, _ := reader.ReadString('\n') // read until the user hits enter

		fmt.Print(info.HeightPromptE)
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
		bmi := (weight / (height * height)) * 703
		// output the calculated BMI
		fmt.Printf("Your BMI: %.2f", bmi)
	}
}
