package controller

import (
    "fmt"
	"time"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/theanik/goapp/app/models"
	"github.com/theanik/goapp/app/initializers"
	"github.com/theanik/goapp/app/utils"
)

func RegisterUser(data map[string]string) {
	projectId, _ := strconv.Atoi(data["projectId"])


	user := models.User{
		Name: data["name"], 
		Email: data["email"], 
		Password: data["password"],
		ProjectId: projectId,
		CreatedAt: time.Now(),
	}
	
	result := initializers.DB.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		fmt.Errorf("User with that email already exists")
		return
	} else if result.Error != nil {
		fmt.Errorf("Something went worng! More info %w", result.Error.Error())
		return
	}
	if result.RowsAffected > 0 {
		fmt.Println("Great! You have beed registard as a user. Your ID is : ", user.ID)
		color.Yellow("Please login with your valid credential now")
	}
}


func Login(data map[string]string) (models.User, error) {

	var user models.User

	result := initializers.DB.First(&user, "email = ?", data["email"])

	if result.Error != nil {
		return user, fmt.Errorf("Invalid credentials")
	}

	if err := utils.VerifyHash(user.Password, data["password"]); err != nil {
		return user, fmt.Errorf("Invalid credentials")
	}

	return user, nil

}