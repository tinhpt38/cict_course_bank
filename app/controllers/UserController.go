package controllers

import (
	"cict-quiz-api/app/database"
	"cict-quiz-api/app/models"
	"context"
	"encoding/json"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) Index() revel.Result  {
	defer c.Request.Destroy()
	var result []models.User
	ctx := context.Background()
	cur, err := database.UserCollection.Find(ctx,bson.D{})
	c.Response.Status = http.StatusInternalServerError
	data := make(map[string]interface{})
	if err != nil {
		data["status"] = "error"
		data["data"] = "Internal Server Error"
		return c.RenderJSON(data)
	}

	for cur.Next(ctx){
		var u models.User
		if err := cur.Decode(&u); err != nil{
			data["status"] = "error"
			data["data"] = "Internal Server Error"
			return c.RenderJSON(data)
		}
		result = append(result,u)
	}

	if err := cur.Err(); err != nil{
		data["status"] = "error"
		data["data"] = "Internal Server Error"
		return c.RenderJSON(data)
	}
	cur.Close(ctx)
	if len(result) == 0{
		data["status"] = "error"
		data["data"] = mongo.ErrNoDocuments
		return c.RenderJSON(data)
	}

	data["status"] = "success"
	data["data"] = result

	return c.RenderJSON(data)
}


func (c UserController) Register() revel.Result{

	defer c.Request.Destroy()
	user := &models.User{}
	err := json.NewDecoder(c.Request.GetBody()).Decode(&user)
	c.Response.Status = http.StatusBadRequest
	data := make(map[string]interface{})
	if err != nil{
		data["status"] = "error"
		data["data"] = "Status Bad Request"
		c.RenderJSON(data)
	}
	hashed ,err := models.Hash(user.Password)
	if err != nil{
		data["status"] = "error"
		data["data"] = "Could hash password"
		c.RenderJSON(data)
	}
	user.Password = hashed
	ctx := context.Background()
	_, err = database.UserCollection.InsertOne(ctx,user)

	if err != nil{
		data["status"] = "error"
		data["data"] = "Could insert user"
	}

	c.Response.Status = http.StatusCreated
	data["status"] = "success"
	data["data"] = user

	return c.RenderJSON(data)
}