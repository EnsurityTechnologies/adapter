package main

import (
	"fmt"

	"github.com/EnsurityTechnologies/adapter"
	"github.com/EnsurityTechnologies/config"
	"golang.org/x/crypto/bcrypt"
)

// UserTable
type UserTable struct {
	UserName     string `gorm:"column:Username;size:128"`
	UserMode     string `gorm:"column:user_mode;size:16"`
	UserPassword string `gorm:"column:user_password;size:128"`
}

func main() {
	config, _ := config.LoadConfig("config.json")
	db, err := adapter.NewAdapter(config)
	if err != nil {
		fmt.Println("Error in new adapter")
		return
	}
	if db.IsTableExist("Tenant", "UserTable") == true {
		db.DropTable("Tenant", "UserTable")
	}
	fmt.Println("Creating Table")
	db.InitTable("Tenant", "UserTable", &UserTable{})
	var userTable UserTable
	userTable.UserMode = "Admin"
	userTable.UserName = "Admin"
	pwdHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	userTable.UserPassword = string(pwdHash)
	db.Create("Tenant", "UserTable", &userTable)
}
