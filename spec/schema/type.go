package schema

type UserType int16

const (
	UserUndefined UserType = iota
	UserUnknown
	UserAdmin
	UserRoot
)

func (u UserType) String() string {
	switch u {
	case UserAdmin:
		return "admin"
	case UserRoot:
		return "root"
	default:
		return "unknown"
	}
}

func UserTypeFromString(s string) UserType {
	switch s {
	case "admin":
		return UserAdmin
	case "root":
		return UserRoot
	default:
		return UserUnknown
	}
}
