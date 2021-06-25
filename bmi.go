package main

import (
	"errors"
	"fmt"
	"runtime"
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

	info.PrintWelcome()

	// prompt for user input (weight + height)
	switch mesurmentType {
	case "metric":
		weight, height, err := getUserMesurments(mesurmentType)
		if err != nil {
			fmt.Println("Something went wrong")
		}
		// bmi := weight / (height * height)
		bmi, err := calculateBMI("metric", weight, height)
		if err != nil {
			fmt.Println("Something went wrong")
		}

		// output the calculated BMI
		fmt.Printf("Your BMI: %.2f", bmi)
	case "english":
		weight, height, err := getUserMesurments(mesurmentType)
		if err != nil {
			fmt.Println("something went wrong")
		}

		// bmi := (weight / (height * height)) * 703
		bmi, err := calculateBMI("english", weight, height)
		if err != nil {
			fmt.Println("Something went wrong")
		}
		// output the calculated BMI
		fmt.Printf("Your BMI: %.2f", bmi)
	}
}

func calculateBMI(mType string, weight, height float64) (float64, error) {
	if mType == "metric" {
		return weight / (height * height), nil
	}
	if mType == "english" {
		return (weight / (height * height)) * 703, nil
	}
	return 0.0, errors.New("Sorry don't support this type")
}
