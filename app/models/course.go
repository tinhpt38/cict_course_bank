package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	CourseName string `bson:"course_name" json:"course_name"`
	Code string `bson:"code" json:"code"`
	Author string `bson:"author" json:"author"`
	Contributors []string `bson:"contributors" json:"contributors"`
	PackagesId []string `bson:"packages_id" json:"packages_id"`
	Image string `bson:"image" json:"image"`
	CreateAt string `bson:"create_at" json:"create_at"`
	DeleteAt string `bson:"delete_at" json:"delete_at"`
	UpdateAt string `bson:"update_at" json:"update_at"`
}

func AddPackageId(course Course, id string)  Course{
	course.PackagesId = append(course.PackagesId, id)
	return course
}

func AddContributors(course Course, id string)  Course{
	course.Contributors = append(course.Contributors, id)
	return course
}

