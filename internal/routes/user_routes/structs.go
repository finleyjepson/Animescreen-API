package user_routes

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsOrganiser bool `json:"isOrganiser"`
	IsAdmin bool `json:"isAdmin"`
	CreatedAt string `json:"createdAt"`
}