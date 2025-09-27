package main

import (
	"github.com/36kone/api-go/database"
	"github.com/36kone/api-go/models"
	"github.com/36kone/api-go/routes"
)

func main() {
	database.DataBaseConnection()
	models.Users = []models.User{}
	routes.HandleRequests()
}
