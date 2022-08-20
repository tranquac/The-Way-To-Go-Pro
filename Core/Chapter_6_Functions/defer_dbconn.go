package main

import "fmt"

func main() {
	dbDBOperation()
}

func connectToDB() {
	fmt.Println("Success, connected to DB!")
}

func disconnectFromDB() {
	fmt.Println("Success, Disconnected from DB!")
}

func dbDBOperation() {
	connectToDB()
	fmt.Println("Defered the database disconnect!")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if there is a return or abnormal termination before
}
