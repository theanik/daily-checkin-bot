package stdform

import (
    "fmt"
	"bufio"

	"github.com/theanik/goapp/app/rules"
	"github.com/theanik/goapp/app/enums"
	"github.com/fatih/color"
)

func CommonRequriedInput(scanner *bufio.Scanner, str *string, label string) (string, bool) {
	isCanceled := false
	for {
		fmt.Println(label)
		fmt.Print("-> ")
		scanner.Scan()
		stdInput := scanner.Text()

		if stdInput == enums.CANCEL {
			color.Blue("Ok, it's cancelled. Once you're ready, please type checkin to fill out your check-in reports.")
			isCanceled = true
			break
		}
		data, err := rules.ValidateRequrid(stdInput, 3)
		if err == nil {
			*str = data
			break
		} else {
			fmt.Println(err)
		}
	}
	return *str, isCanceled
}