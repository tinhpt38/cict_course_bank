package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	Name string `bson:"name" json:"name"`
	Code string `bson:"code" json:"code"`
	Author string `bson:"author" json:"author"`
	Contributors []string `bson:"contributors" json:"contributors"`
	PackagesId []string `bson:"packages_id" json:"packages_id"`
	Image string `bson:"image" json:"image"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	DeletedAt string `bson:"deleted_at" json:"deleted_at"`
	UpdatedAt string `bson:"updated_at" json:"updated_at"`
}

//func AddPackageId(course Course, id string)  Course{
//	course.PackagesId = append(course.PackagesId, id)
//	return course
//}
//
//func AddContributors(course Course, id string)  Course{
//	course.Contributors = append(course.Contributors, id)
//	return course
//}

