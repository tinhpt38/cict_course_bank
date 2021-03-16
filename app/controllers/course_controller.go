package controllers

import (
	"cict-quiz-api/app/database"
	"cict-quiz-api/app/models"
	"context"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type CourseController struct {
	*revel.Controller
}

func (c CourseController) Index() revel.Result {
	defer c.Request.Destroy()
	result := []models.Course{}
	ctx := context.Background()
	cur, err := database.UserCollection.Find(ctx, bson.D{})
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

	data["status"] = "success"
	data["data"] = result

	return c.RenderJSON(data)
}