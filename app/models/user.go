package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FullName string `bson:"full_name" json:"full_name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email string `bson:"email" json:"email"`
	RoleID string `bson:"role_id" json:"role_id"`
	TokenID string `bson:"token_id" json:"token_id"`
	Avatar string 	`bson:"avatar" json:"avatar"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	DeletedAt string `bson:"deleted_at" json:"deleted_at"`
	UpdatedAt string `bson:"updated_at" json:"updated_at"`
}

func Hash(password string)(string,error){
	bytes, err := bcrypt.GenerateFromPassword( []byte(password),14)
	return string(bytes),err

}

func CheckHashAndPassword(password, hash string) error  {
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
}