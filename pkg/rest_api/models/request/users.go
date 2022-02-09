package request

type User struct {
	UserID    int    `json:"user_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	DOB       string `json:"DOB" binding:"required" time_format:"2006-01-02"`
	CreatedOn int64  `json:"created_on" binding:"required"`
}

type UserRequestBody []User
