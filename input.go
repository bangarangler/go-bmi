package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/bangarangler/go-bmi/info"
)

// can't be a const compile don't run functions
// even though once function runs it won't change again that function is only
// ran at runtime
var reader = bufio.NewReader(os.Stdin)

func getUserMesurments(mType string) (float64, float64, error) {
	switch mType {
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

		return weight, height, nil
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
		return weight, height, nil
	default:
		return 0.0, 0.0, errors.New("Unsupported type")
	}
}
