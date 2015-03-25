package controllers

import (
	"github.com/revel/revel"
	"chickspace/app/models"
	//"chickspace/app"
)

type Home struct {
	*revel.Controller
}

func (c Home) Index() revel.Result {
	//db, _ := sqlx.Connect("mysql", "root:aaaaaa@/chickspace")
	//defer db.Close()
	
	//db.MustExec(`INSERT INTO accounts VALUES (null, ?, ?, ?)`, "nightwolfz", "aaaaaa", "ryan@megidov.com")
	//db.MustExec(`INSERT INTO profiles VALUES (null, ?)`, "nightwolfz")
	//inserted, _ := result.LastInsertId()
	
	/*accounts := []models.Account{}
	err := app.DB.Select(&accounts, "SELECT id,username,password,email FROM accounts")
	if err != nil {
		revel.ERROR.Println("DB Error", err)
	}*/
	
	account := models.Account{}.Get("nightwolfz")
	
	
	
	/*accounts := make(map[string]map[string]string)
	rows, err := db.Query("SELECT username, password FROM accounts")
	for rows.Next() {
		var username string
		var password string
		err = rows.Scan(&username, &password)
		if err != nil {
			fmt.Print(err)
		}
		accounts[username] = map[string]string{
			"username": username,
			"password": password,
		}
	}*/
	
	return c.Render(account)
}
