package models

import (
	"github.com/revel/revel"
	"chickspace/app"
)

type Profile struct {
	Id      int
    Name 	string
}

func (t Profile) Get(id int) Profile {
	profile := Profile{}
	err := app.DB.Get(&profile, "SELECT id, name FROM profiles WHERE id=? LIMIT 1", id)
	if err != nil {
		revel.ERROR.Print(err)
	}
	revel.INFO.Print(profile)
	
	return profile
}
/*
func (t *Profile) Validate(v *revel.Validation) {

	v.Check(t.Name,
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
	
}}*/