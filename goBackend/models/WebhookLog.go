package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoWebhookLog struct {
	Id                primitive.ObjectID  `bson:"_id"`
	Timestamp         primitive.Timestamp `bson:"timestamp"`
	WebhookId         primitive.ObjectID  `bson:"webhookId"`
	ForwardSuccessful bool                `bson:"forwardSuccessful"`
	Data              bson.M              `bson:"data"`
}

type WebhookLog struct {
	Timestamp         primitive.Timestamp `bson:"timestamp"`
	WebhookId         primitive.ObjectID  `bson:"webhookId"`
	ForwardSuccessful bool                `bson:"forwardSuccessful"`
	Data              bson.M              `bson:"data"`
}
