package actions

import (
	"fmt"
	"bufio"
	"strconv"

	"github.com/fatih/color"
	"github.com/theanik/goapp/app/enums"
	"github.com/theanik/goapp/app/utils"
	"github.com/theanik/goapp/app/stdform"
	"github.com/theanik/goapp/app/controller"
	"github.com/theanik/goapp/app/models"
)

func CheckAuth(scanner *bufio.Scanner, authUser *models.User) {
	color.Cyan("Please login with you valid credentials!")
	for {
		color.Blue("If you alreay have an account type \"login\" otherwise type \"register\"")
		fmt.Print(" -> ")
		scanner.Scan()
		text := scanner.Text()

		if text == enums.LOGIN {
			user, err := loginAction(scanner)
			if err != nil {
				fmt.Println(err)
				color.Red("Invalid credentials")
			} else {
				*authUser = user
				color.HiGreen("Great! Welcome %s", user.Name)
				break
			}
		} else if text == enums.REGISTER {
			registerAction(scanner)
		}
	}
}


/*
*/
func loginAction(scanner *bufio.Scanner) (models.User, error) {
	email := ""
	password := ""
	stdform.EmailInput(scanner, &email)
	stdform.PaswordInput(scanner, &password)

	loginPayload := map[string]string{
		"email": email,
		"password": password,
	}
	user, err := controller.Login(loginPayload)
	return user, err
}


func registerAction(scanner *bufio.Scanner) {
	name := ""
	email := ""
	password := ""
	stdform.NameInput(scanner, &name)
	stdform.EmailInput(scanner, &email)
	stdform.PaswordInput(scanner, &password)
	passwordHash, _ := utils.MakeHash(password)
	projectId := stdform.ProjectInput(scanner)

	userData := map[string]string{
		"name": name,
		"email": email,
		"password": passwordHash,
		"projectId": strconv.Itoa(projectId),
	}

	controller.RegisterUser(userData)
}
