package services

import (
	"context"
	"crud_gin_gonic/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServicee struct {
	usercollection *mongo.Collection
	ctx context.Context
}

func NewUserService(usercollection *mongo.Collection,ctx context.Context) UserService{
	return &UserServicee {
		usercollection: usercollection,
		ctx: ctx,
	}
}


func (u *UserServicee) CreateUser(user *models.User) error{
	_, err := u.usercollection.InsertOne(u.ctx,user)
	return err
}


func (u *UserServicee) GetUser(name *string)(*models.User,error){
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx,query).Decode(&user)
	return user,err 
}

func (u *UserServicee) GetAll() ([]*models.User,error) {
	var users []*models.User
	cursor , err := u.usercollection.Find(u.ctx,bson.D{{}})
	if err!= nil {
		return nil,err
	}
	for cursor.Next(u.ctx){
		var user models.User
		err:= cursor.Decode(&user)

		if err!= nil {
			return nil,err
		}
		users= append(users, &user)
	}
	if err:= cursor.Err(); err!=nil {
		return nil,err
	}

	cursor.Close(u.ctx)

	if len(users)== 0 {
		return nil,errors.New("Document Not Found")
	}
	return users,nil
}


func (u *UserServicee) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key:"user_name", Value: user.Name}, bson.E{Key:"user_age", Value: user.Age},bson.E{Key:"user_address", Value: user.Add}}}}
	result,_ := u.usercollection.UpdateOne(u.ctx,filter,update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}


func (u *UserServicee) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value : name}}
	result,_:= u.usercollection.DeleteOne(u.ctx, filter)

	if result.DeletedCount !=1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}