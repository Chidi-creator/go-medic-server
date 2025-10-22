package utils

func IsRoleValid(roles []Roles, userRoles []Roles) bool {
	for _, userRole := range userRoles {
		for _, validRole := range roles {
			if userRole == validRole {
				return true
			}
		}
	}
	return false
}
