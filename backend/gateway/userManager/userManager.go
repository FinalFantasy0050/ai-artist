package usermanager

func Authentication(user string) bool {
	_, exist := usermap[user]

	return exist
}
