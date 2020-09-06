package migrations

import (
	"github.com/jessicaramsa/fintech-banking-app/database"
	"github.com/jessicaramsa/fintech-banking-app/helpers"
	"github.com/jessicaramsa/fintech-banking-app/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {
	users := &[2]interfaces.User{
		{Username: "Jessica", Email: "jessicaramsa@gmail.com"},
		{Username: "Jenkins", Email: "jensskins@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{
			Type:    "Daily Account",
			Name:    string(users[i].Username + "'s account"),
			Balance: uint(10000 * int(i+1)),
			UserID:  user.ID}
		database.DB.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)
	createAccounts()
}
