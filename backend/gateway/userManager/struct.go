package usermanager

const fileName = "userManager/user.json"

var usermap map[string]bool

type users struct {
	User []string `json:"user"`
}
