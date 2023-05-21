package controller

import (
    "fmt"
	"time"
	"strconv"
	// "gorm.io/gorm"
	// "strings"

	"github.com/fatih/color"
	"github.com/theanik/goapp/app/models"
	"github.com/theanik/goapp/app/initializers"
	// "github.com/theanik/goapp/app/utils"
)

// var db *gorm.DB


// func init() {
// 	initializers.ConnectDB()
// 	// db = initializers.GetDB()
// }

func CreateCheckinRecord(data map[string]string) {
	userId, _ := strconv.Atoi(data["userId"])
	projectId, _ := strconv.Atoi(data["projectId"])
	fmt.Println(data)
	report := models.DailyReport{
		TodayDescription: data["today"], 
		PreviousDayDescription: data["previousDay"], 
		Blocker: data["blocker"],
		UserId: userId,
		ProjectId: projectId,
		CreatedAt: time.Now(),
	}

	result := initializers.DB.Create(&report)

	if result.Error != nil {
		color.Red("Some went wrong!")
		fmt.Errorf("More info %w", result.Error.Error())
		return
	}
	if result.RowsAffected > 0 {
		msg := "Well done! This is all, you can continue with your work 💪"
		color.Yellow(msg)
	}
}

func CheckinReport(data map[string]string) {
	var reports []models.DailyReport
	today := data["date"] + " 00:00:00"
	todayLast := data["date"] + " 23:59:59"
	initializers.DB.Preload("User").Where("created_at BETWEEN ? AND ?", today, todayLast).Find(&reports)
	
	fmt.Println("")
	fmt.Println("Date", data["date"])

	// lavel := color.New(color.FgYellow).SprintFunc()
	//  := color.New(color.FgCyan).SprintFunc()
	if len(reports) > 0 {
		for _, report := range reports {
			fmt.Println("")
			color.Cyan("----------------------------------------------------------------------")
			fmt.Println(color.YellowString("User : "), report.User.Name)
			fmt.Println(color.YellowString("Today Plan : "), report.TodayDescription)
			fmt.Println(color.YellowString("Previous Day Tasks : "), report.PreviousDayDescription)
			fmt.Println(color.YellowString("Blocker : "), report.Blocker)
			color.Cyan("----------------------------------------------------------------------")
			fmt.Println("")
		}
	} else {
		fmt.Println("")
		color.Cyan("----------------------------------------------------------------------")
		fmt.Println(color.RedString("No data found 😢"))
		color.Cyan("----------------------------------------------------------------------")
		fmt.Println("")
	}

	
}