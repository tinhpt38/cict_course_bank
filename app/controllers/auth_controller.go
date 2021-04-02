package controllers

import (
	"cict-quiz-api/app/database"
	"cict-quiz-api/app/models"
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

//login
type AuthController struct {
	* revel.Controller
}

func (c AuthController) Index() revel.Result{
	return revel.RenderHTMLResult{};
}

func (c AuthController) Login() revel.Result {
	defer c.Request.Destroy()
	login := &models.Login{}
	c.Response.Status = http.StatusBadRequest
	data := make(map[string]interface{})
	data["status"] = "error"
	if err := json.NewDecoder(c.Request.GetBody()).Decode(&login); err != nil{
		data["data"] = "Could parse request"
		return c.RenderJSON(data)
	}

	if govalidator.IsNull(login.Username) || govalidator.IsNull(login.Password){
		data["data"] = "Username or Password is not correct"
		return c.RenderJSON(data)
	}

	result := &models.User{}
	ctx := context.Background()
	filter := bson.D{primitive.E{Key: "username", Value: login.Username}}

	if err := database.UserCollection.FindOne(ctx,filter).Decode(&result); err != nil{
		data["data"] = "Username not exist"
		return c.RenderJSON(data)
	}

	if err := models.CheckHashAndPassword(login.Password,result.Password); err !=nil{
		data["data"] = "Username or Password is not correct"
		return c.RenderJSON(data)
	}

	// create token
	// token, errCreate := jwt.Create(username)
	// if errCreate != nil {
	// 	c.Response.Status = http.StatusBadRequest
	// 	var body interface{}
	// 	body = "t"
	// 	message := &models.Message{Name:"err",Body:body}
	// 	return c.RenderJSON(message)
	// }

	c.Response.Status = http.StatusOK
	body := make(map[string]interface{})
	body["user"] = result
	body["token"] = ""
	data["data"] = body
	return c.RenderJSON(data)
}