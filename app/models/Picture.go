package models

import (
	"github.com/revel/revel"
	"chickspace/app"
)

type Picture struct {
	Id      int
    AccountId 	string
    Src 	string
}

func (t Picture) Get(accountId int) []Picture {
	pictures := []Picture{}
	err := app.DB.Get(&pictures, "SELECT id, accountId, src FROM pictures WHERE accountId='?'", accountId)
	if err != nil {
		revel.ERROR.Print(err)
	}
	return pictures
}



func (t *Picture) Validate(v *revel.Validation) {

	/*v.Check(t.Name,
		revel.Required{},
		revel.MaxSize{40},
		revel.MinSize{4},
	)

	v.Check(t.Name, revel.MaxSize{40})
	v.Check(t.Job, revel.MaxSize{40})
	v.Check(t.Reference, revel.MaxSize{40})

	if t.Phone == "" {
		v.Check(t.Email,
			revel.Required{},
			revel.MaxSize{20},
			revel.MinSize{4},
		)
	}

	if t.Email == "" {
		v.Check(t.Phone,
			revel.Required{},
			revel.MaxSize{20},
			revel.MinSize{4},
		)
	}*/
}