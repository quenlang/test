package main

import (
	"fmt"
)

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to db")
	return
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
	return
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return
}
