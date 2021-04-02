package controllers

import "github.com/revel/revel"

//questions
type QuestionController struct {
	*revel.Controller
}


func(c *QuestionController) Index() revel.Result{
	return c.RenderText("Hello")
}