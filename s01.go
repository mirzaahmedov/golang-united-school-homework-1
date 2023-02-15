package structs

type User struct {
	firstName string
	lastName string
}

func New(firstName, lastName string) *User {
	return &User{
		firstName: firstName,
		lastName: lastName,
	}
}
func ResetUser(user *User) {
	user.firstName = ""
	user.lastName = ""
}
func IsUser(value interface{}) bool {
	switch value.(type) {
	case User:
		return true
	default:
		return false
	}
}
func ProccessUser(user UserInterface) string {
	user.SetFirstName("John")
	user.SetLastName("Doe")

	return user.FullName()
}

func (u *User)SetFirstName(value string) {
	u.firstName = value
}
func (u *User)SetLastName(value string) {
	u.lastName = value
}
func (u *User)FullName() string {
	return u.lastName + " " + u.firstName
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}
