package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"

  "po_users/general"
)

func updateUserDB(userID string, user general.User) (error) {
  collection := client.Database(dataBaseName).Collection(collectionName)

  filter := bson.D{{"id", userID}}
  update := bson.D{{"$set", user}}

  _, err := collection.UpdateOne(context.TODO(), filter, update)
  if err != nil {
    return err
  }

  return nil
}
