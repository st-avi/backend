package consts

type Role int

const (
	RoleUser  Role = 1
	RoleAdmin Role = 2
)

func (r Role) String() string {
	switch r {
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	default:
		return ""
	}
}
