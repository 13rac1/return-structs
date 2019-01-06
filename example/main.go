package main

type UserInterface interface {
	GetName() string
}

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func getUser(name string) UserInterface {
	return &User{Name: name}
}

func getUsers(name string) interface{} {
	return &User{Name: name}
}

func putUser(user UserInterface) {

}

func main() {
	u := getUser("bob")

	if u.GetName() == "bob" {
		// Comment.
	}
}
