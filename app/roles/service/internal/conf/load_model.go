package conf

type SuperUser struct {
	SuperUser []SuperUserItem `json:"super_user"`
}

type SuperUserItem struct {
	Name string
	Pass string
}
