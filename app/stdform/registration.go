package stdform

import (
    "fmt"
	"bufio"

	"github.com/theanik/goapp/app/rules"
)

func NameInput(scanner *bufio.Scanner, name *string) string {
	for {
		fmt.Println("Enter your name : ")
		fmt.Print(" -> ")
		scanner.Scan()
		inputName := scanner.Text()
		validate_name, err := rules.ValidateRequrid(inputName, 3)
		if err == nil {
			*name = validate_name
			break
		} else {
			fmt.Println(err)
		}
	}
	return *name
}

func EmailInput(scanner *bufio.Scanner, email *string) string {
	for {
		fmt.Println("Enter your email : ")
		fmt.Print(" -> ")
		scanner.Scan()
		stdData := scanner.Text()
		validate_data, err := rules.ValidateEmail(stdData)
		if err == nil {
			*email = validate_data
			break
		} else {
			fmt.Println(err)
		}
	}
	return *email
}

func PaswordInput(scanner *bufio.Scanner, str *string) string {
	for {
		fmt.Println("Enter your password : ")
		fmt.Print(" -> ")
		scanner.Scan()
		stdData := scanner.Text()
		validate_data, err := rules.ValidateRequrid(stdData, 6)
		if err == nil {
			*str = validate_data
			break
		} else {
			fmt.Println(err)
		}
	}
	return *str
}


func ProjectInput(scanner *bufio.Scanner) int {
	for {
		fmt.Println("Enter your projectId : ")
		fmt.Print(" -> ")
		scanner.Scan()
		stdData := scanner.Text()
		validate_data, err := rules.ValidateInteger(stdData)
		if err == nil {
			return validate_data
			break
		} else {
			fmt.Println(err)
		}
	}
	return 0
}
