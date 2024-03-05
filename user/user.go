package user

type User struct {
	Name  string
	Email string
}

var users []User

func Register(u User) {
	users = append(users, u)
}

func GetUser(email string) (User, bool) {
	for _, u := range users {
		if u.Email == email {
			return u, true
		}
	}
	return User{}, false
}
