package request

type User struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	DOB       string `json:"DOB"`
	CreatedOn int    `json:"created_on"`
}
