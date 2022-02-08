package response

type User struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	DOBDayOfWeek string `json:"DOB_day_of_week"`
	CreatedOn    string `json:"created_on"`
}

type UserResponse []User
