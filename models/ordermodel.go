package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_date time.Time          `bson:"order_date" validate:"required"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"updated_at"`
	Order_id   string             `bson:"order_id"`
	Table_id   *string            `bson:"table_id" validate:"required"`
}
