package playground

type User struct {
	Id       int64  `json:"id,omitempty"`
	UniqueId string `json:"unique_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`

	Salt     string `json:"salt,omitempty"`
	Password string `json:"password,omitempty"`
}
