package actions

import (
	"fmt"
	"bufio"
	"strconv"
	"time"
	"regexp"

	"github.com/fatih/color"
	"github.com/theanik/goapp/app/stdform"
	"github.com/theanik/goapp/app/enums"
	"github.com/theanik/goapp/app/controller"
	"github.com/theanik/goapp/app/models"
)

func CheckinAction(scanner *bufio.Scanner, authUser models.User) {
	previousDayStd := ""
	todayDayStd := ""
	hasBlockerStd := ""
	color.Yellow(`Type "cancel" to descard the progress`)

	label := "What did you complete in your previous workday?"
	_, cancel := stdform.CommonRequriedInput(scanner, &previousDayStd, label)
	if cancel == true {
		return
	}

	label = "What are you planning to work on today?"
	_, cancel = stdform.CommonRequriedInput(scanner, &todayDayStd, label)
	if cancel == true {
		return
	}

	label = "Great. Do you have any blockers? If so, just tell me. Otherwise please say: \"no\"."
	_, cancel = stdform.CommonRequriedInput(scanner, &hasBlockerStd, label)
	if cancel == true {
		return
	}

	projectIdStr := strconv.Itoa(authUser.ProjectId)
	userIdInt := int(authUser.ID)
	userIdStr := strconv.Itoa(userIdInt)
	reportData := map[string]string{
		"today": todayDayStd,
		"previousDay": previousDayStd,
		"blocker": hasBlockerStd,
		"projectId": projectIdStr,
		"userId": userIdStr,
	}

	controller.CreateCheckinRecord(reportData)
}

func ReportAction(str string) {

	t := time.Now()
	reportTime := t.Format("2006-01-02")

	if str != enums.REPORT {
		inStr := str

		re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	
		if re.MatchString(inStr) {
			dateStr := re.FindAllString(inStr, -1)[0]
			isValid := isDateValid(dateStr)
			if isValid == false {
				color.Red("Given date format is not valid. Format eg : \"report 2023-05-17\"")
				return
			}
			reportTime = dateStr
	
		} else {
			color.Red("Given date format is not valid. Format eg : \"report 2023-05-17\"")
			return
		}
	}

	reportDataParams := map[string]string{
		"date": reportTime,
	}

	fmt.Println(reportDataParams)

	controller.CheckinReport(reportDataParams)
}

func isDateValid(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
 }