package database

import (
  "po_users/general"
)

func UpdateUser(user general.User) {
  userID := user.ID
  _, err := FindUser_ID(userID)
  if err != nil {
    insertUser(user)
  }else {
    updateUserDB(userID, user)
  }
}
