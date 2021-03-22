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

type CategoryController struct {
	*revel.Controller
}

func (c CategoryController) Index() revel.Result {

	return c.RenderText("Index")
}

func (c CategoryController) GetAll() revel.Result {
	defer c.Request.Destroy()
	var result = []models.Category{}
	ctx := context.Background()
	cur, err := database.CategoryCollection.Find(ctx, bson.D{})
	c.Response.Status = http.StatusInternalServerError
	data:= make(map[string]interface{})
	if err != nil {
		data["status"] = "error"
		data["data"] = "Internal Server Error"
		return c.RenderJSON(data)
	}

	for cur.Next(ctx){
		var cat models.Category
		if err:= cur.Decode(&cat); err !=nil{
			data["status"] = "error"
			data["data"] = "Internal Server Error"
			return c.RenderJSON(data)
		}
		result = append(result, cat)
	}
	if err := cur.Err(); err != nil {
		data["status"] = "error"
		data["data"] = "Internal Server Error"
		return c.RenderJSON(data)
	}
	cur.Close(ctx)
	if len(result) == 0 {
		data["status"] = "error"
		data["data"] = mongo.ErrNoDocuments
		return c.RenderJSON(data)
	}
	c.Response.Status = http.StatusOK
	data["status"] = "success"
	data["data"] = result

	return c.RenderJSON(data)
}

func (c CategoryController) Create() revel.Result {
	defer c.Request.Destroy()
	var cat = &models.Category{}
	err := json.NewDecoder(c.Request.GetBody()).Decode(&cat)
	c.Response.Status = http.StatusBadRequest
	data := make(map[string]interface{})
	ctx := context.Background()
	data["status"] = "Error"

	if err != nil{
		data["data"] = "Status Bad Request. Could not decode Category"
		c.RenderJSON(data)
	}

	_, err = database.CategoryCollection.InsertOne(ctx, cat)
	ctx.Done()
	if err !=nil{
		data["status"] = "error"
		data["data"] = "Could insert user"
		return c.RenderJSON(data)
	}

	c.Response.Status = http.StatusCreated
	data["status"] = "success"
	data["data"] = cat

	return c.RenderJSON(data)

}