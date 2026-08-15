package auth

import "fmt"


// if you want to access this function outside of this package then you need to start function name with capital letter

func LoginWithCredentials(username string, password string){
	fmt.Println("login user using", username, password)
}