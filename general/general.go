package general

type Project struct {
  ID string
  Name string
}

type User struct {
  ID        string `json:"id"`
  SessionID string `json:"sessionid"`
  Name      string `json:"name"`
  Projects  []Project `json:"projects"`

  GoogleID  string `json:"googleid"`
}
