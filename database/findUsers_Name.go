package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo/options"

  "po_users/general"
)

func FindUsers_Name(name string) ([]general.User, error) {
  collection := client.Database(dataBaseName).Collection(collectionName)

  find := bson.D{{"name", name}}

  cur, err := collection.Find(context.TODO(), find, options.Find())
  if err != nil {
    return nil, err
  }

  result := make([]general.User, 0)
  for cur.Next(context.TODO()) {
    var elem general.User
    err := cur.Decode(&elem)
    if err != nil {
      continue
    }

    result = append(result, elem)
  }

  cur.Close(context.TODO())

  return result, nil
}
