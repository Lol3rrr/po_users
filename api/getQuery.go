package api

import (
  "net/http"
)

func getQuery(r *http.Request) (map[string][]string) {
  return r.URL.Query()
}
