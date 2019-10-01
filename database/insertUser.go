package database

import (
  "context"

  "po_users/general"
)

func insertUser(user general.User) (error) {
  collection := client.Database(dataBaseName).Collection(collectionName)

  _, err := collection.InsertOne(context.TODO(), user)
  if err != nil {
    return err
  }

  return nil
}
