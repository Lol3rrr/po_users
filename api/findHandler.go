package api

import (
  "net/http"
  "encoding/json"

  "po_users/general"
  "po_users/database"
)

type FindResponse struct {
  Results []general.User `json:"results"`
}

func sendResult(resp FindResponse, w http.ResponseWriter) {
  jsonResponse, err := json.Marshal(resp)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}

func findWithID(id string, w http.ResponseWriter) {
  res, err := database.FindUser_ID(id)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  results := make([]general.User, 1)
  results[0] = res

  response := FindResponse{
    Results: results,
  }

  sendResult(response, w)
}

func findWithSessionID(id string, w http.ResponseWriter) {
  res, err := database.FindUser_sessionID(id)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  results := make([]general.User, 1)
  results[0] = res

  response := FindResponse{
    Results: results,
  }

  sendResult(response, w)
}

func findWithName(name string, w http.ResponseWriter) {
  res, err := database.FindUsers_Name(name)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  response := FindResponse{
    Results: res,
  }

  sendResult(response, w)
}

func findHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  id, id_ok := query["id"]
  if id_ok && len(id) > 0 {
    findWithID(id[0], w)
    return
  }

  sessionID, sessionID_ok := query["sessionID"]
  if sessionID_ok && len(sessionID) > 0 {
    findWithSessionID(sessionID[0], w)
    return
  }

  name, name_ok := query["name"]
  if name_ok && len(name) > 0 {
    findWithName(name[0], w)
    return
  }


  w.WriteHeader(200)
}
