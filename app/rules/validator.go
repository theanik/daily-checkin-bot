package rules

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"net/mail"
)



func ValidateRequrid(str string, min int)  (string, error){
	if len(str) >= min {
		return str, nil
	}
	red := color.New(color.FgRed).PrintfFunc()
	red("This text is requried with %d minimum chars", min)
	fmt.Println("")
	return "", fmt.Errorf("")
}

func ValidateEmail(str string)  (string, error){
	_, err := mail.ParseAddress(str)
	if err == nil{
		return str, nil
	}
	red := color.New(color.FgRed).PrintlnFunc()
	red("Please enter a valid mail")
    return "", fmt.Errorf("Error: %s", err)
}

func ValidateInteger(str string) (int, error) {
	if number, err := strconv.Atoi(str); err == nil {
		return number, nil
	}
	red := color.New(color.FgRed).PrintlnFunc()
	red("Please enter a valid integer, eg : 5")
	return 0, fmt.Errorf("")
}
