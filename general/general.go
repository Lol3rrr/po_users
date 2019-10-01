package general

type Project struct {
  ID string
  Name string
}

type User struct {
  ID        string
  SessionID string
  Name      string
  Projects  []Project
}
