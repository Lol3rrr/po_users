package general

type Project struct {
  ID string
  Name string
}

type User struct {
  ID        string    `json:"id"`
  SessionID string    `json:"sessionid"`
  Name      string    `json:"name"`
  Projects  []Project `json:"projects"`

  GoogleID  string    `json:"googleid"`
}

type ResponseUser struct {
  ID        string    `json:"id"`
  Name      string    `json:"name"`
  Projects  []Project `json:"projects"`
}

func (user *User) ConvertToResponse() (ResponseUser) {
  result := ResponseUser{
    ID: user.ID,
    Name: user.Name,
    Projects: user.Projects,
  }

  return result
}
