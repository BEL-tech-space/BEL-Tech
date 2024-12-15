package pages

type Collector struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Role      string `json:"role"`
}
