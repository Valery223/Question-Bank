package domain

// UserRole - роль пользователя в системе для авторизации
type UserRole string

const (
	RoleNone    UserRole = "none"
	RoleAdmin   UserRole = "admin"
	RoleManager UserRole = "manager"
	RoleUser    UserRole = "user"
	RoleGuest   UserRole = "guest"
)

func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleManager, RoleUser, RoleGuest:
		return true
	}
	return false
}

func (r UserRole) CanCreateQuestions() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanDeleteQuestions() bool {
	return r == RoleAdmin
}

func (r UserRole) CanViewQuestions() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanUpdateQuestions() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanCreateTemplates() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanDeleteTemplates() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanViewTemplates() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanUpdateTemplates() bool {
	return r == RoleAdmin || r == RoleManager
}

func (r UserRole) CanCreateSessions() bool {
	return r == RoleAdmin || r == RoleManager || r == RoleUser
}

func (r UserRole) CanViewSessions() bool {
	return r == RoleAdmin || r == RoleManager || r == RoleUser
}

func (r UserRole) CanViewAllSessions() bool {
	return r == RoleAdmin || r == RoleManager
}
