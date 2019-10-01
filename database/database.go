package database

import (
  "go.mongodb.org/mongo-driver/mongo"
)

var dbURL string

var client *mongo.Client

const dataBaseName = "users"
const collectionName = "users"
