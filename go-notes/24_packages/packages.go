package main

import (
	"fmt"
	"github.com/fatih/color"

	"github.com/karangupta982/golang/user"


	"github.com/karangupta982/golang/auth"
)

// for creating package we need to initialize a module, write in terminal (go mod init your_project_name) generally giturl or project is given

// login functionality is reused multiple time so we can create package of that functionality

// to install other packages in go then go to https://pkg.go.dev/ and search for package then go get github.com/fatih/color

// main func is the entry point of your program
func main(){

	auth.LoginWithCredentials("codersgyan", "secret")


	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Name: "Karan gupta",
		Email: "user@gmail.com",
	}
	// fmt.Println(user)
	color.Red(user.Name)

}
