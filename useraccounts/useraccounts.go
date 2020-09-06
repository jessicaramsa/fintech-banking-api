package useraccounts

import (
	"fmt"

	"github.com/jessicaramsa/fintech-banking-app/database"
	"github.com/jessicaramsa/fintech-banking-app/helpers"
	"github.com/jessicaramsa/fintech-banking-app/interfaces"
	"github.com/jessicaramsa/fintech-banking-app/transactions"
)

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ? ", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAccount := interfaces.ResponseAccount{}

	database.DB.Where("id = ? ", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAccount.ID = account.ID
	responseAccount.Name = account.Name
	responseAccount.Balance = int(account.Balance)
	return responseAccount
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)
		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You are not owner of the account"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is too small"}
		}
		
		updatedAccount := updateAccount(from, int(fromAccount.Balance) - amount)
		updateAccount(to, int(toAccount.Balance) + amount)

		transactions.CreateTransaction(from, to, amount)
		
		var response = map[string]interface{}{"message": "OK"}
		response["data"] = updatedAccount
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
