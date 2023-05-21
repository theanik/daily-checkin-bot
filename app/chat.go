package main

import (
	"bufio"
    "fmt"
    "os"
	"strings"

	"github.com/fatih/color"
	"github.com/theanik/goapp/app/enums"
	"github.com/theanik/goapp/app/models"
	"github.com/theanik/goapp/app/actions"
	"github.com/theanik/goapp/app/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var authUser models.User

	

	if authUser == (models.User{}) {
		actions.CheckAuth(scanner, &authUser)
	}
	
	color.Cyan(`Type "help" for keyword list`)
    for {
        fmt.Print(" -> ")
        scanner.Scan()
        text := scanner.Text()


		// Report input
		if strings.HasPrefix(text, "report") {
			actions.ReportAction(text)
		}

		// Common keywords
		switch text {
		case enums.HELP:
			color.Cyan(enums.HELP_RESPONSE)
		case enums.CHECKIN:
			actions.CheckinAction(scanner, authUser)
		}

		


		
		if text == enums.LOGOUT {
			fmt.Println("Bye! Have a good day")
			break
		}
    }
}


