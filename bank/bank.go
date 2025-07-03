package main

import "fmt"

func main() {
    fmt.Println("Welcome to Go bank!")
    fmt.Println("What do you want to do?")
    fmt.Println("1. Check balance")
    fmt.Println("2. Deposit money")
    fmt.Println("3. Withdraw money")
    fmt.Println("4. Exit")
    
    var userChoice int
    fmt.Print("Your choice: ")
    fmt.Scan(&userChoice)
    
    fmt.Println("Your choice:", userChoice)
}