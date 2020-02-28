//Nicholas Larsen
//February 25, 2020
//User plays computer in Rock, Paper, Scissors
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var user int
	rand.Seed(time.Now().UnixNano())
	computer := (rand.Intn(3) + 1)
	//uses time and rand to have computer be a random number between 1 and 3

	fmt.Println("Enter 1 for Rock, 2 for Scissors, and 3 for Paper")
	fmt.Scanln(&user)
	//User's choice

	if computer == 1 {
		fmt.Println("computer chooses Rock")

	} else if computer == 2 {
		fmt.Println("computer chooses Scissors")

	} else {
		fmt.Println("computer chooses Paper")
		//tells the computer's choice
	}
	if user == 1 {
		fmt.Println("Player chooses Rock")

	} else if user == 2 {
		fmt.Println("User chooses Scissors")

	} else {
		fmt.Println("User chooses Paper")
		//tells the User's choice
  }
  if computer==2 && user==3{
  fmt.Println("You have lost!")

}else if computer==1 && user==2{
  fmt.Println("You have lost!")

}else if computer==3 && user==1{
  fmt.Println("You have lost!")

}else if computer==user{
    fmt.Println("It is a Tie")

}else {
  fmt.Println("YOU HAVE WON")
  //adds more else if statements to tell who wins or not
}
}
