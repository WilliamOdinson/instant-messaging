package main

import (
	"fmt"
)

var userId int
var userPwd string

func main() {
	// Receive user’s choice
	var key int

	// Control whether to continue showing the menu
	var loop = true

	for loop {
		fmt.Println("-- Welcome to the Multi-User Chat System --")
		fmt.Println("1. Login to Chatroom")
		fmt.Println("2. Register Account")
		fmt.Println("3. Exit System")
		fmt.Println("Please enter your choice (1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("[System] Logging into the chatroom...")
		case 2:
			fmt.Println("[System] Redirecting to register a new account...")
		case 3:
			fmt.Println("[System] Bye bye!")
			loop = false
		default:
			fmt.Println("Invalid input, please enter a number between 1 and 3.")
		}
		if key == 1 {
			fmt.Println("[System] Enter the id of the user")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("[System] Enter the password of the user")
			fmt.Scanf("%s\n", &userPwd)
			err := login(userId, userPwd)
			if err != nil {
				fmt.Println("Login failed")
			} else {
				fmt.Println("Login succeeded.")
			}
		}
	}

}
