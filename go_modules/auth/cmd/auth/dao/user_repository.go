package dao

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/model"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const USER_COLLECTION_NAME = "user"
const STUDENT_ROLE = "STUDENT"
const TEACHER_ROLE = "TEACHER"
const ADMIN_ROLE = "ADMIN"

func GetUserByTelegramChatId(chatId int) model.User {
	var result model.User
	filter := bson.M{"telegramchatid": chatId}

	getUserCollection().FindOne(context.TODO(), filter).Decode(&result)

	return result
}

func CreateUser(user model.User) model.User {
	user.Roles = []string{STUDENT_ROLE}

	bytes, _ := json.Marshal(user)
	var unmarshaledUser model.User
	var jsonUser string
	json.Unmarshal(bytes, &unmarshaledUser)
	json.Unmarshal(bytes, &jsonUser)

	getUserCollection().InsertOne(context.TODO(), user)

	return user
}

func getUserCollection() *mongo.Collection {
	return config.DbManager.Collection(USER_COLLECTION_NAME)
}
