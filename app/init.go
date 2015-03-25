package app

import (
	"github.com/revel/revel"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

// GLOBALS
var DB *sqlx.DB


func InitDB() {
	
	connectionString := revel.Config.StringDefault("db.user", "") + ":" + revel.Config.StringDefault("db.pass", "")
	DB, _ = sqlx.Connect("mysql", connectionString + "@/" + revel.Config.StringDefault("db.database", ""))
	//defer DB.Close()
	revel.INFO.Println("DB Connected")
	
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	
	revel.OnAppStart(InitDB)
	//revel.InterceptMethod((*Database).Connect, revel.BEFORE)
	//revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

/*
func (c Hotels) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}
	return nil
}*/


// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
