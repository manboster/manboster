package types

type UserType int16

const (
	UserUnknown UserType = iota
	UserRoot    UserType = iota
	UserAdmin   UserType = iota
)

type User struct {
	ID uint64
}
