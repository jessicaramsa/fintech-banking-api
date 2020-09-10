package main

import (
	"github.com/jessicaramsa/fintech-banking-app/api"
	"github.com/jessicaramsa/fintech-banking-app/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
