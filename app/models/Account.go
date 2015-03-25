package models

import (
	"github.com/revel/revel"
	"chickspace/app"
)

type Account struct {
	Id       int
    Username string
	Password string
	Email string
}

func (t Account) Get(username string) Account {
	account := Account{}
	err := app.DB.Get(&account, "SELECT * FROM accounts WHERE username=? LIMIT 1", username)
	if err != nil {
		revel.ERROR.Print(err)
	}
	revel.INFO.Print(account)
	return account
}