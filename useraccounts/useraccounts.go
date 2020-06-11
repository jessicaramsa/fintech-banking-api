package useraccounts

import (
	"github.com/jessicaramsa/fintech-banking-app/helpers"
	"github.com/jessicaramsa/fintech-banking-app/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ? ", id).Update("balance", ammount)
	defer db.Close()
}
