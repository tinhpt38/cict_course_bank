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

type CourseController struct {
	*revel.Controller
}
//course
func (c CourseController) Index() revel.Result {
	return c.RenderText("Inndex")
}

func (c CourseController) GetAll() revel.Result {
	defer c.Request.Destroy()
	result := []models.Course{}
	ctx := context.Background()
	cur, err := database.CourseCollection.Find(ctx, bson.D{})
	c.Response.Status = http.StatusInternalServerError
	data := make(map[string]interface{})
	if err != nil {
		data["status"] = "error"
		data["data"] = "Internal Server Error"
		return c.RenderJSON(data)
	}

	for cur.Next(ctx) {
		var u models.Course
		if err := cur.Decode(&u); err != nil {
			data["status"] = "error"
			data["data"] = "Internal Server Error"
			return c.RenderJSON(data)
		}
		result = append(result, u)
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

func (c CourseController) InsertOne() revel.Result{
	defer c.Request.Destroy()
	course := &models.Course{}
	c.Response.Status = http.StatusBadRequest
	data := make(map[string]interface{})
	err := json.NewDecoder(c.Request.GetBody()).Decode(&course)
	data["status"] = "error"
	if err != nil{
	data["data"] = "status Bad Request " + err.Error()
	return c.RenderJSON(data)
	}
	ctx := context.Background()
	_, err = database.CourseCollection.InsertOne(ctx, course)
	if err != nil {
		data["status"] = "error"
		data["data"] = "Could insert user"
		return c.RenderJSON(data)
	}
	ctx.Done()
	c.Response.Status = http.StatusCreated
	data["status"] = "success"
	data["data"] = course

	return c.RenderJSON(data)
}

func (c CourseController) InsertMany() revel.Result{
	defer c.Request.Destroy()
	courses := []models.Course{}
	err := json.NewDecoder(c.Request.GetBody()).Decode(&courses)
	c.Response.Status = http.StatusBadRequest
	data := make(map[string]interface{})
	data["status"] = "error"
	if err != nil{
		data["data"] = "status Bad Request"
	}
	ctx := context.Background()
	documents := make([]interface{},len(courses))
	documents = append(documents, courses)
	_, err = database.CourseCollection.InsertMany(ctx, documents)
	if err != nil {
		data["status"] = "error"
		data["data"] = "Could insert user"
		return c.RenderJSON(data)
	}
	c.Response.Status = http.StatusCreated
	data["status"] = "success"
	data["data"] = documents

	return c.RenderJSON(data)
}
