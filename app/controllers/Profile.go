package controllers

import (
	"github.com/revel/revel"
	"chickspace/app/models"
)


type Profile struct {
	*revel.Controller
}

func (c Profile) Index() revel.Result {

	profile := models.Profile{}.Get(1)
	pictures := models.Picture{}.Get(profile.Id)

	return c.Render(profile, pictures)
}