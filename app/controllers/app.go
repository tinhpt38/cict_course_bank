package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}
//home
func (c App) Index() revel.Result {
	return c.Render()
}
