package info

import "fmt"

const (
	MesurmentType = "Type metric for (kg) or english for (lb): "
	mainTitle     = "BMI Calculator"
	seperator     = "--------------------"
	// Sane rest of world
	WeightPromptM = "Please enter your weight (kg): "
	HeightPromptM = "Please enter your height (m): "
	// US
	WeightPromptE = "Please enter your weight (lb): "
	HeightPromptE = "Please enter your height (in): "
)

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(seperator)
}
