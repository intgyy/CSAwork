package dao

var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

func Adduser(username, Password string) {
	database[username] = Password
}
func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}
func SelectPasswordFromUsername(username string) string {
	return database[username]
}
