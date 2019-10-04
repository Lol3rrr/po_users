package api

func getQueryElement(query map[string][]string, key string) (string, bool) {
  rawValue, found := query[key]
  if !found || len(rawValue) <= 0 {
    return "", false
  }

  return rawValue[0], true
}
