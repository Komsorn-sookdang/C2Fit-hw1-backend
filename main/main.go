package main

import (
	"c2fit-hw-backend/databases"
	"c2fit-hw-backend/router"
	"c2fit-hw-backend/session"
)

func main() {
	databases.ConnectToDatabase()

	session.CreateStore()

	r := router.SetUpRouter()

	r.Run(":8080")
}
