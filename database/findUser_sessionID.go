package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"

  "po_users/general"
)

func FindUser_sessionID(sessionID string) (general.User, error) {
  collection := client.Database(dataBaseName).Collection(collectionName)
  var result general.User

  filter := bson.D{{"sessionid", sessionID}}

  err := collection.FindOne(context.TODO(), filter).Decode(&result)
  if err != nil {
    return result, err
  }

  return result, nil
}
