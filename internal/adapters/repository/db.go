package repository

import (
	"context"

	"github.com/immersivesky/go-rest-mongo/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	userCollection *mongo.Collection
}

func NewDB(dsn string) *DB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	return &DB{
		userCollection: client.Database("vk").Collection("chats"),
	}
}

func (db *DB) GetChat(chatID int) *domain.Chat {
	chat := new(domain.Chat)

	if err := db.userCollection.FindOne(context.Background(), bson.D{{"chat_id", chatID}}).Decode(&chat); err != nil {
		return nil
	}

	return chat
}

func (db *DB) CreateChat(chatID int) bool {
	result, err := db.userCollection.InsertOne(context.Background(), bson.D{{"chat_id", chatID}})
	if err != nil {
		panic(err)
	}

	if result.InsertedID != nil {
		return true
	}

	return false
}
