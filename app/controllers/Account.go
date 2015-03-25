package controllers

import (
	"github.com/revel/revel"
	"chickspace/app/models"
    "chickspace/app/routes"
	"chickspace/app"
)

type Account struct {
	*revel.Controller
}

func (c Account) Login(post models.LoginPost) revel.Result {

	account := models.Account{}
	err := app.DB.Get(&account, "SELECT id,username,password,email FROM accounts WHERE email='?' AND password='?' LIMIT 1", post.Username, post.Password)
	if err != nil {
		revel.ERROR.Print(err)
	}
	
    return c.Redirect(routes.Profile.Index())
}