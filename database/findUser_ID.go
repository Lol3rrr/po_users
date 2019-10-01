package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"

  "po_users/general"
)

func FindUser_ID(userID string) (general.User, error) {
  collection := client.Database(dataBaseName).Collection(collectionName)
  var result general.User

  filter := bson.D{{"id", userID}}

  err := collection.FindOne(context.TODO(), filter).Decode(&result)
  if err != nil {
    return result, err
  }

  return result, nil
}
