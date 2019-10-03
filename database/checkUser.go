package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"

  "po_users/general"
)

func UserExists(user general.User) (bool, general.User) {
  if user.GoogleID != "" {
    result, user := existsRaw("googleid", user.GoogleID)

    if result {
      return true, user
    }
  }

  return false, general.User{}
}

func existsRaw(key, value string) (bool, general.User) {
  collection := client.Database(dataBaseName).Collection(collectionName)
  var result general.User

  filter := bson.D{{key, value}}

  err := collection.FindOne(context.TODO(), filter).Decode(&result)
  if err != nil {
    return false, general.User{}
  }

  return true, result
}
