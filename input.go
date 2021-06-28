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
		weight := getUserInput(info.WeightPromptM, mType)
		height := getUserInput(info.HeightPromptM, mType)

		return weight, height, nil
	case "english":
		weight := getUserInput(info.WeightPromptE, mType)
		height := getUserInput(info.HeightPromptE, mType)

		return weight, height, nil
	default:
		return 0.0, 0.0, errors.New("Unsupported type")
	}
}

func getUserInput(promptText, mType string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n') // read until the user hits enter

	switch mType {
	case "english":
		if runtime.GOOS == "windows" {
			userInput = strings.Replace(userInput, "\r\n", "", -1)
		} else {
			userInput = strings.Replace(userInput, "\n", "", -1)
		}
	case "metric":
		if runtime.GOOS == "windows" {
			userInput = strings.Replace(userInput, "\r\n", "", -1)
		} else {
			userInput = strings.Replace(userInput, "\n", "", -1)
		}
	}

	enteredValue, _ := strconv.ParseFloat(userInput, 64)
	return enteredValue
}
