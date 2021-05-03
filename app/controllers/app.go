package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}


func (c App) LoginUser(username,password string) revel.Result {

	if username=="ThuyPhuong"&& password=="1234" {
		c.Flash.Success("Dang nhap thanh cong")
		return c.Redirect(App.Question)
	}

	c.Flash.Error("Thong tin dang nhap khong hop le")
	return c.Redirect(App.Index)
	//return c.Render(username)
}
func (c App) Question() revel.Result  {
	return c.Render()
}


