package main

import (
	"fmt"
	"github.com/theanik/goapp/app/initializers"
	"github.com/theanik/goapp/app/models"

)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? User Migration complete")
}